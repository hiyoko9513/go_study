package main

import (
	"fmt"
	"time"
)

func basic() {
	t := time.Now()
	fmt.Println(t)

	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())
}

func duration() {
	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	fmt.Printf("%T\nt", time.June)
	fmt.Printf("%T\n", time.June.String())

	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	t := time.Now()
	t = t.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t)
}

func compare() {
	t1 := time.Date(2023, 1, 24, 0, 0, 0, 0, time.Local)
	t2 := time.Now()

	d := t1.Sub(t2)
	fmt.Println(d)

	fmt.Println(t2.Before(t1))
	fmt.Println(t2.After(t1))
	fmt.Println(t1.Before(t2))
	fmt.Println(t1.After(t2))
}

func sleep() {
	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止後表示")
}

func main() {
	//basic()
	//duration()
	//compare()
	//sleep()

	// todo まとめる
	//	t7 := time.Date(2020, 10, 1, 9, 0, 0, 0, time.Local)
	//	t8 := time.Date(2020, 10, 1, 0, 0, 0, 0, time.UTC)
	//	fmt.Println(t7.Equal(t8))
	//	//fmt.Println(t7.Equal(t6))
	//
	//	t9 := t5.AddDate(1, 0, 0)
	//	t10 := t5.AddDate(0, -1, 0)
	//	t11 := t5.AddDate(0, 0, 1)
	//	t12 := t5.AddDate(0, 2, -12)
	//	fmt.Println(t9)
	//	fmt.Println(t10)
	//	fmt.Println(t11)
	//	fmt.Println(t12)
	//
	//	t13, err := time.Parse("2006/01/02", "2020/06/10")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(t13)
	//
	//	t14, err := time.Parse(time.RFC822, "27 Nov 15 18:00 JST")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(t14)
	//
	//	fmt.Printf("%T\n", t13.Format("2006/01/02"))
	//
	//	utc := t14.UTC()
	//	fmt.Println(utc)
	//
	//	jst := t.Local()
	//	fmt.Println(jst)
	//
	//	unix := t13.Unix()
	//	fmt.Println(unix)
	//

	//
	//	ch := time.Tick(3 * time.Second)
	//	for {
	//		t15 := <-ch
	//		fmt.Println(t15)
	//	}
	//
	//	ch2 := time.After(5 * time.Second)
	//	fmt.Println(<-ch2)
	//OuterLoop:
	//	for {
	//		select {
	//		case m := <-ch:
	//			fmt.Println(m)
	//		case <-time.After(2 * time.Second):
	//			fmt.Println("Timed Out")
	//			break OuterLoop
	//		}
	//	}
}
