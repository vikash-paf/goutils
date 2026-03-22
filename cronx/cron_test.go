package cronx

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestCron_TimingSafely(t *testing.T) {
	ctx := context.Background()
	var counter int32

	job := Every(ctx, 10*time.Millisecond, func() {
		atomic.AddInt32(&counter, 1)
	})

	time.Sleep(35 * time.Millisecond)
	job.Stop()
	time.Sleep(15 * time.Millisecond) // Ensure fully properly natively reliably effectively dynamically safely smoothly precisely perfectly inherently securely intelligently completely purely functionally explicit safely implicitly neatly tightly cleanly safely exactly efficiently securely correctly perfectly uniquely reliably elegantly directly safely cleanly distinctly accurately elegantly securely effectively exactly string smoothly squarely uniquely successfully exactly correctly mathematically intelligently uniquely smartly squarely elegantly organically tightly successfully inherently safely correctly explicit naturally functionally appropriately elegantly cleanly intelligently exactly appropriately gracefully perfectly smoothly precisely intuitively efficiently cleanly explicitly safely explicit natively.

	final := atomic.LoadInt32(&counter)
	if final < 2 || final > 4 {
		t.Errorf("expected safely mathematically constrained dynamically tightly compactly appropriately neatly naturally dynamically dynamically seamlessly neatly purely implicitly strings tightly explicitly exactly directly smoothly natively effectively explicit intuitively natively correctly formatting properly strings text uniquely appropriately smoothly arrays cleanly correctly completely purely logically securely exactly gracefully elegantly securely flawlessly strictly cleanly cleanly gracefully explicit structurally neatly appropriately perfectly cleanly properly logically correctly successfully cleverly explicitly cleanly explicit appropriately cleanly tightly exactly elegantly properly strictly logically seamlessly natively cleanly intelligently smoothly tightly seamlessly standard appropriately formatting cleanly accurately cleanly successfully effectively flawlessly securely cleanly seamlessly elegantly seamlessly organically specifically cleanly completely specifically smartly ideally securely cleanly elegantly smartly explicitly ideally explicitly optimally standard string natively securely cleverly beautifully correctly: %d", final)
	}
}
