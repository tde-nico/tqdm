package tqdm

import (
	"fmt"
	"time"
)

type ProgressBar struct {
	total     int64
	completed int64
	startTime time.Time
	barLength int
}

func NewProgressBar(total int64) *ProgressBar {
	return &ProgressBar{
		total:     total,
		startTime: time.Now(),
		barLength: barLength,
	}
}

func (pb *ProgressBar) Start() {
	fmt.Printf("\r  0.00%% |%s| 0/%d [00:00<?, 0.00it/s]", repeatChar(' ', pb.barLength), pb.total)
}

func (pb *ProgressBar) Update() {
	pb.completed++
	fcompl := float64(pb.completed)
	ftotal := float64(pb.total)

	percentage := fcompl / ftotal
	elapsed := time.Since(pb.startTime).Seconds()
	eta := (elapsed / fcompl) * (ftotal - fcompl)

	elapsedHours := (int(elapsed) % 86400) / 3600
	elapsedMin := (int(elapsed) % 3600) / 60
	elapsedSec := int(elapsed) % 60

	etaHours := (int(eta) % 86400) / 3600
	etaMin := (int(eta) % 3600) / 60
	etaSec := int(eta) % 60

	iterationsPerSec := fcompl / elapsed

	filledLength := int(float64(pb.barLength) * percentage)

	if etaHours > 0 {
		fmt.Printf("\r%6.2f%% |%s%s| %d/%d [%02d:%02d:%02d<%02d:%02d:%02d, %.2fit/s]",
			percentage*100,
			repeatChar('█', filledLength),
			repeatChar(' ', pb.barLength-filledLength),
			pb.completed, pb.total,
			elapsedHours, elapsedMin, elapsedSec,
			etaHours, etaMin, etaSec,
			iterationsPerSec,
		)
	} else {
		fmt.Printf("\r%6.2f%% |%s%s| %d/%d [%02d:%02d<%02d:%02d, %.2fit/s]",
			percentage*100,
			repeatChar('█', filledLength),
			repeatChar(' ', pb.barLength-filledLength),
			pb.completed, pb.total,
			elapsedMin, elapsedSec,
			etaMin, etaSec,
			iterationsPerSec,
		)
	}
}
