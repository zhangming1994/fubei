package snowflake

import (
	"testing"
)

func TestUID(t *testing.T) {
	max := 10000

	iw, _ := NewIDWorker(1)

	uidsChan := make(chan int64)

	for i := 0; i < max; i++ {
		go func(iw *IDWorker, uidsChan chan int64) {
			id, err := iw.NextID()
			if err != nil {
				t.Fatal(err)
			} else {
				uidsChan <- id
			}
		}(iw, uidsChan)
	}

	uids := []int64{}
	for i := 0; i < max; i++ {
		uids = append(uids, <-uidsChan)
	}

	if len(uids) != max {
		t.Error("生成数据失败")
	}

	//uidsUniq := sliceutil.UniqI64Slice(uids)
	//if len(uidsUniq) != max {
	//	t.Errorf("没有生成唯一的ID! len: %d", len(uidsUniq))
	//}

	//fmt.Println(uidsUniq)

}
