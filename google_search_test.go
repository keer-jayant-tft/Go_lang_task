package main

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/chromedp/chromedp"
)

var _ = Describe("Google Search Automation", func() {
	var ctx context.Context
	var cancel context.CancelFunc
	var allocCancel context.CancelFunc

	BeforeEach(func() {
		// Step 1: Create an ExecAllocator with visible Chrome
		var allocCtx context.Context
		allocCtx, allocCancel = chromedp.NewExecAllocator(context.Background(),
			append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", true),        // run in visible mode
				chromedp.Flag("disable-gpu", false),    // use GPU
				chromedp.Flag("start-maximized", true), // full window
			)...,
		)

		// Step 2: Create a new browser context
		ctx, cancel = chromedp.NewContext(allocCtx)
	})

	AfterEach(func() {
		cancel()
		allocCancel()
	})

	It("should search for 'Apple device' and click the first link", func() {
		var firstLink string
		err := chromedp.Run(ctx,

			// i was facing catcha issue with  google chrome so i try to seach with duckducgo

			// chromedp.Navigate(`https://www.google.com`),
			// chromedp.Sleep(2*time.Second), //

			// chromedp.WaitVisible(`#APjFqb`, chromedp.ByID),
			// chromedp.SendKeys(`#APjFqb`, "Apple device\n", chromedp.ByID),
			// chromedp.WaitVisible(`#search`, chromedp.ByID),

			// chromedp.Sleep(2*time.Second),
			// chromedp.Click(`h3`, chromedp.NodeVisible, chromedp.ByQuery), // first result
			// chromedp.Sleep(3*time.Second),

			chromedp.Navigate(`https://duckduckgo.com/`),
			chromedp.Sleep(2*time.Second),

			chromedp.SendKeys(`#searchbox_input`, "Apple device\n", chromedp.ByID),
			chromedp.WaitVisible(`[data-testid="result-title-a"]`, chromedp.ByQuery),
			chromedp.Sleep(3*time.Second),
			chromedp.Text(`(//a[@data-testid="result-title-a"])[2]`, &firstLink, chromedp.BySearch),
			chromedp.Click(`(//a[@data-testid="result-title-a"])[2]`, chromedp.BySearch),

			chromedp.Sleep(7*time.Second),
		)

		Expect(err).To(BeNil())
	})
})
