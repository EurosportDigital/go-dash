package mpd

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func makePSSHBox(systemID, payload []byte) ([]byte, error) {
	psshBuf := &bytes.Buffer{}
	size := uint32(12 + 16 + 4 + len(payload)) // 3 ints, systemID, "pssh" string and payload

	if len(systemID) != 16 {
		return nil, fmt.Errorf("SystemID must be 16 bytes, was: %d", len(systemID))
	}

	if err := binary.Write(psshBuf, binary.BigEndian, size); err != nil {
		return nil, err
	}

	if err := binary.Write(psshBuf, binary.BigEndian, []byte("pssh")); err != nil {
		return nil, err
	}

	if err := binary.Write(psshBuf, binary.BigEndian, uint32(0)); err != nil {
		return nil, err
	}

	if _, err := psshBuf.Write(systemID); err != nil {
		return nil, err
	}

	if err := binary.Write(psshBuf, binary.BigEndian, uint32(len(payload))); err != nil {
		return nil, err
	}

	if _, err := psshBuf.Write(payload); err != nil {
		return nil, err
	}

	return nil, nil
}
