package main

import (
	"fmt"

	"github.com/zencoder/go-dash/mpd"
)

func main() {

	m := mpd.NewMPD(mpd.DASH_PROFILE_LIVE, "PT6M16S", "PT1.97S")

	audioAS, _ := m.AddNewAdaptationSetAudio(mpd.DASH_MIME_TYPE_AUDIO_MP4, true, 1, "und")

	audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	audioAS.AddNewContentProtectionSchemeWidevine()
	audioAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	audioAS.SetNewSegmentTemplate(1968, "$RepresentationID$/audio/en/init.mp4", "$RepresentationID$/audio/en/seg-$Number$.m4f", 0, 1000)
	audioAS.AddNewRepresentationAudio(44100, 67095, "mp4a.40.2", "800")

	videoAS, _ := m.AddNewAdaptationSetVideo(mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)

	videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	videoAS.AddNewContentProtectionSchemeWidevine()
	videoAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	videoAS.SetNewSegmentTemplate(1968, "$RepresentationID$/video/1/init.mp4", "$RepresentationID$/video/1/seg-$Number$.m4f", 0, 1000)
	videoAS.AddNewRepresentationVideo(1518664, "avc1.4d401f", "800", "30000/1001", 960, 540)
	videoAS.AddNewRepresentationVideo(1911775, "avc1.4d401f", "1000", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2295158, "avc1.4d401f", "1200", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2780732, "avc1.4d401f", "1500", "30000/1001", 1280, 720)

	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(mpd.DASH_MIME_TYPE_SUBTITLE_VTT, "en")
	subtitleRep, _ := subtitleAS.AddNewRepresentationSubtitle(256, "subtitle_en")
	subtitleRep.SetNewBaseURL("http://example.com/content/sintel/subtitles/subtitles_en.vtt")

	mpdStr, _ := m.WriteToString()
	fmt.Println(mpdStr)
}
