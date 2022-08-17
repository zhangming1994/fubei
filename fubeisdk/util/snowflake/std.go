package snowflake

import "time"

var std *IDWorker

func Init(workerID int64) (err error) {
	std, err = NewIDWorker(workerID)
	return
}

func NextID() int64 {
	id, err := std.NextID()
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		id, _ = std.NextID()
	}
	return id
}
