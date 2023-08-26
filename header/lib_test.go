package header_test

import (
	"bytes"
	"testing"

	"github.com/GoNZooo/newsboat-helpers/header"
)

func TestParse(t *testing.T) {
	data := []byte(`Feed: Oxide and Friends
Title: No Silver Bullets
Author: Oxide Computer Company
Date: Tue, 15 Aug 2023 03:00:00 +0300
Link: https://share.transistor.fm/s/d584546b
Podcast Download URL: https://media.transistor.fm/d584546b/dbe949dc.mp3 (type: audio/mpeg)

Bryan and Steve Klabnik discuss Fred Brooks' essay "No Silver
Bullets"--ostensibly apropos of nothing!--discussing the challenges to 10x (or
100x!) improvements in software engineering.

In addition to Bryan Cantrill[1] speakers on included Steve Klabnik[2], Ian
Grunert[3], and Tom Lyon[4].

Some of the topics we hit on, in the order that we hit them:

  * No Silver Bullet[5] by Fred Brooks
  * Sub-podcasting (it's a thing!) this[6]
  * video: Fred Brooks speaking on No Silver Bullet[7]
  * Ruby on Rails demo[8] (2005)
  * Future of coding podcast[9]
  * Amdahl's law[10]
  * FizzBuzzEnterpriseEdition[11]
  * Knuth and McIlroy Approach a Problem[12]

If we got something wrong or missed something, please file a PR! Our next show
will likely be on Monday at 5p Pacific Time on our Discord server; stay tuned to
our Mastodon feeds for details, or subscribe to this calendar[13]. We'd love to
have you join us, as we always love to hear from new speakers!

Links:
[1]: https://mastodon.social/@bcantrill (link)
[2]: https://twitter.com/steveklabnik (link)
[3]: https://hachyderm.io/@iangrunert (link)
[4]: https://mastodon.social/@aka_pugs (link)
[5]: http://worrydream.com/refs/Brooks-NoSilverBullet.pdf (link)
[6]: https://redplanetlabs.com/ (link)
[7]: https://www.youtube.com/watch?v=HWYrrw7Zf1k (link)
[8]: https://www.youtube.com/watch?v=Gzj723LkRJY (link)
[9]: https://futureofcoding.org/ (link)
[10]: https://en.wikipedia.org/wiki/Amdahl%27s_law (link)
[11]: https://github.com/EnterpriseQualityCoding/FizzBuzzEnterpriseEdition (link)
[12]: https://matt-rickard.com/instinct-and-culture (link)
[13]: https://sesh.fyi/api/calendar/v2/iMdFbuFRupMwuTiwvXswNU.ics (link)
`)
	r := bytes.NewReader(data)
	header, err := header.Parse(r)
	if err != nil {
		t.Errorf("Error parsing header: %v", err)
	}

	if header.Feed != "Oxide and Friends" {
		t.Errorf("Unexpected feed: %v", header.Feed)
	}

	if header.Title != "No Silver Bullets" {
		t.Errorf("Unexpected title: %v", header.Title)
	}

	if header.Author != "Oxide Computer Company" {
		t.Errorf("Unexpected author: %v", header.Author)
	}

	if header.Link != "https://share.transistor.fm/s/d584546b" {
		t.Errorf("Unexpected link: %v", header.Link)
	}

	if header.PodcastUrl != "https://media.transistor.fm/d584546b/dbe949dc.mp3" {
		t.Errorf("Unexpected podcast URL: %v", header.PodcastUrl)
	}
}
