package rss

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestItem(t *testing.T) {
	blob := `
      <item>
        <title>A Beginner&#39;s Look at BenchmarkTools.jl</title>
        <link>https://randyzwitch.com/benchmarktools-julia-benchmarking/</link>
        <guid isPermaLink="false">https://lobste.rs/s/m9zan1</guid>
        <author>jstuartmill@users.lobste.rs (jstuartmill)</author>
        <pubDate>Fri, 20 Dec 2019 21:51:22 -0600</pubDate>
        <comments>https://lobste.rs/s/m9zan1/beginner_s_look_at_benchmarktools_jl</comments>
        <description>
          
            &#60;p&#62;&#60;a href=&#34;https://lobste.rs/s/m9zan1/beginner_s_look_at_benchmarktools_jl&#34;&#62;Comments&#60;/a&#62;&#60;/p&#62;
        </description>
          <category>programming</category>
      </item>`

	got := Item{}
	err := xml.Unmarshal([]byte(blob), &got)
	if err != nil {
		t.Errorf("%q", err)
	}
	fmt.Println(got)
	fmt.Println(time.Time(got.PubDate))

	if len(got.Categories) != 1 {
		t.Errorf("len(Categories) = %d; want 1", len(got.Categories))
	}
}

func TestRss(t *testing.T) {
	file, err := os.Open("test/programming.rss")
	if err != nil {
		return
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	rss := Rss{}
	decoder.Decode(&rss)
	//fmt.Println(rss)
	if len(rss.Channel.Items) == 0 {
		fmt.Println("oops")
	}
}
