package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveRadio struct {
	Slug    string
	Name    string
	Owner   string
	URL     string
	Cards   []string
	Links   []socialLink
	AddedAt time.Time
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		{
			Slug:  "lofigirl",
			Name:  "Lofi Girl",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/jfKfPfyJRdk",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is shown studying in Lyon, a city from France where her designer Juan Pablo used to live.",
				"The view through the window depicts the buildings on the slopes of Croix-Rousse, where the bell tower of the Bon-Pasteur church can be spotted.",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://lofigirl.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/jfKfPfyJRdk",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/lofigirl",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/LofiGirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/lofigirl",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/lofigirl",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/lofigirl",
				},
			},
		},
		{
			Slug:  "lofigirl-sleepy",
			Name:  "Lofi Sleepy Girl",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/rUxyKA_-grg",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is living in Lyon, a city from France where her designer Juan Pablo used to live.",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://lofigirl.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/rUxyKA_-grg",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/lofigirl",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/LofiGirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/lofigirl",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/lofigirl",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/lofigirl",
				},
			},
		},
		{
			Slug:  "chillhop",
			Name:  "Chillhop Raccoon",
			Owner: "Chillhop Music",
			URL:   "https://youtu.be/5yx6BWlEVcY",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d?si=3co_7Q6XT0OJZwkBlqWoDQ",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:  "chillhop-relax",
			Name:  "Chillhop Relaxing Raccoon",
			Owner: "Chillhop Music",
			URL:   "https://youtu.be/7NOSDKb0HlU",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/7NOSDKb0HlU",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d?si=3co_7Q6XT0OJZwkBlqWoDQ",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:  "collegemusic-guy",
			Name:  "College Guy",
			Owner: "College Music",
			URL:   "https://youtu.be/QwXHcgZUnFI",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/QwXHcgZUnFI",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/collegemusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/collegemusic",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/collegemusicyt",
				},
			},
		},
		{
			Slug:  "lofi-code-beats",
			Name:  "Coding Beats",
			Owner: "Joma Tech",
			URL:   "https://youtu.be/PY8f1Z3nARo",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/PY8f1Z3nARo",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/jomaoppa",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/jomaoppa",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/jomaoppa",
				},
			},
		},
		{
			Slug:  "dreamhop",
			Name:  "Dreamhop",
			Owner: "Dreamhop Music",
			URL:   "https://youtu.be/wkhLHTmS_GI",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://dreamhopmusic.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/wkhLHTmS_GI",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/user/91jfqzlv0htrqrvvc60138qma",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/FxF9kng",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/dreamhopp",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Dreamhopp",
				},
			},
		},
		{
			Slug:  "synthwave-radio",
			Name:  "Synthwave Radio",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/4xDzrJKXOOY",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "bootleg",
			Name:  "Bootleg",
			Owner: "the bootleg boy",
			URL:   "https://youtu.be/A_hmrykwR7g",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "peaceful-piano",
			Name:  "Peaceful Piano",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/nWjC1RnLYbM",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "dark-ambience",
			Name:  "Dark Ambient Radio",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/SKhpl1OMqEY",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "tavern",
			Name:  "Tavern Radio",
			Owner: "Filip Lackovic",
			URL:   "https://youtu.be/vK5VwVyxkbI",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "fantasy",
			Name:  "Fantasy Radio",
			Owner: "Michael Ghelfi Studio",
			URL:   "https://youtu.be/DEJNkrHCUTE",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "shaaaaark",
			Name:  "Silly Shark Playlist",
			Owner: "dolly life",
			URL:   "https://youtu.be/7FEhvFgBSq4",
			Cards: []string{},
			Links: []socialLink{},
		},
	}

	sort.Slice(radios, func(a, b int) bool {
		return strings.Compare(radios[a].Name, radios[b].Name) < 0
	})

	for _, r := range radios {
		sort.Slice(r.Links, func(a, b int) bool {
			return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
		})
	}

	return radios
}

type socialLink struct {
	Slug string
	URL  string
}

func socialIcon(slug string) string {
	switch slug {
	case "youtube":
		return youtubeSVG

	case "reddit":
		return redditSVG

	case "facebook":
		return facebookSVG

	case "instagram":
		return instagramSVG

	case "twitter":
		return twitterSVG

	case "spotify":
		return spotifySVG

	case "discord":
		return discordSVG

	case "website":
		return websiteSVG

	default:
		return linkSVG
	}
}
