package snowflake

import (
	"errors"
	"net"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *snowflake
)

const (
	timestampBits = 41
	machineBits   = 8
	sequenceBits  = 14
	maxMachineID  = -1 ^ (-1 << machineBits)
	maxSequence   = -1 ^ (-1 << sequenceBits)
	timeShift     = machineBits + sequenceBits
	machineShift  = sequenceBits
	startEpoch    = int64(1700000000000) // 自定义起始时间戳（2023-11-15）
)

type snowflake struct {
	mu        sync.Mutex
	lastStamp int64
	sequence  int64
	machineID int64
}

func initInstance() {
	mid, _ := getMachineID()
	instance = &snowflake{
		machineID: mid,
	}
}

func getMachineID() (int64, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return 0, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip := ipnet.IP.To4(); ip != nil {
				return int64(ip[3]), nil
			}
		}
	}
	return 0, errors.New("no valid IP found")
}

func GenID() (uint64, error) {
	once.Do(initInstance)
	id, err := instance.nextID()
	return uint64(id), err
}

func (s *snowflake) nextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()

	// 时钟回拨保护
	if now < s.lastStamp {
		offset := s.lastStamp - now
		if offset > 5000 {
			return 0, errors.New("clock moved backwards over 5 seconds")
		}
		time.Sleep(time.Duration(offset) * time.Millisecond)
		now = time.Now().UnixMilli()
	}

	if now == s.lastStamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for now <= s.lastStamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = now

	return ((now - startEpoch) << timeShift) |
		(s.machineID << machineShift) |
		s.sequence, nil
}
