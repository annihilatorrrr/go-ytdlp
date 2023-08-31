// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Post-Processing Option Group

package ytdlp

type PostProcessingBuilder struct {
	parent *Command
}

// Then jumps back to the base command builder, if you want to add additional flags
// from another flag builder.
func (ff *PostProcessingBuilder) Then() *Command {
	return ff.parent
}

// Convert video files to audio-only files (requires ffmpeg and ffprobe)
//
// ExtractAudio maps to cli flags: -x/--extract-audio.
func (ff *PostProcessingBuilder) ExtractAudio() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "extractaudio",
		Flag: "--extract-audio",
		Args: nil,
	})
	return ff
}

// Format to convert the audio to when -x is used. (currently supported: best
// (default), aac, alac, flac, m4a, mp3, opus, vorbis, wav). You can specify
// multiple rules using similar syntax as --remux-video
//
// AudioFormat maps to cli flags: --audio-format=FORMAT.
func (ff *PostProcessingBuilder) AudioFormat(format string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "audioformat",
		Flag: "--audio-format",
		Args: []string{format},
	})
	return ff
}

// Specify ffmpeg audio quality to use when converting the audio with -x. Insert a
// value between 0 (best) and 10 (worst) for VBR or a specific bitrate like 128K
// (default %default)
//
// AudioQuality maps to cli flags: --audio-quality=QUALITY.
func (ff *PostProcessingBuilder) AudioQuality(quality string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "audioquality",
		Flag: "--audio-quality",
		Args: []string{quality},
	})
	return ff
}

// Remux the video into another container if necessary (currently supported: avi,
// flv, gif, mkv, mov, mp4, webm, aac, aiff, alac, flac, m4a, mka, mp3, ogg, opus,
// vorbis, wav). If target container does not support the video/audio codec,
// remuxing will fail. You can specify multiple rules; e.g. "aac>m4a/mov>mp4/mkv"
// will remux aac to m4a, mov to mp4 and anything else to mkv
//
// RemuxVideo maps to cli flags: --remux-video=FORMAT.
func (ff *PostProcessingBuilder) RemuxVideo(format string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "remuxvideo",
		Flag: "--remux-video",
		Args: []string{format},
	})
	return ff
}

// Re-encode the video into another format if necessary. The syntax and supported
// formats are the same as --remux-video
//
// RecodeVideo maps to cli flags: --recode-video=FORMAT.
func (ff *PostProcessingBuilder) RecodeVideo(format string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "recodevideo",
		Flag: "--recode-video",
		Args: []string{format},
	})
	return ff
}

// Keep the intermediate video file on disk after post-processing
//
// KeepVideo maps to cli flags: -k/--keep-video.
func (ff *PostProcessingBuilder) KeepVideo() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "keepvideo",
		Flag: "--keep-video",
		Args: nil,
	})
	return ff
}

// Delete the intermediate video file after post-processing (default)
//
// NoKeepVideo maps to cli flags: --no-keep-video.
func (ff *PostProcessingBuilder) NoKeepVideo() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "keepvideo",
		Flag: "--no-keep-video",
		Args: nil,
	})
	return ff
}

// Overwrite post-processed files (default)
//
// PostOverwrites maps to cli flags: --post-overwrites.
func (ff *PostProcessingBuilder) PostOverwrites() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "nopostoverwrites",
		Flag: "--post-overwrites",
		Args: nil,
	})
	return ff
}

// Do not overwrite post-processed files
//
// NoPostOverwrites maps to cli flags: --no-post-overwrites.
func (ff *PostProcessingBuilder) NoPostOverwrites() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "nopostoverwrites",
		Flag: "--no-post-overwrites",
		Args: nil,
	})
	return ff
}

// Embed subtitles in the video (only for mp4, webm and mkv videos)
//
// EmbedSubs maps to cli flags: --embed-subs.
func (ff *PostProcessingBuilder) EmbedSubs() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "embedsubtitles",
		Flag: "--embed-subs",
		Args: nil,
	})
	return ff
}

// Do not embed subtitles (default)
//
// NoEmbedSubs maps to cli flags: --no-embed-subs.
func (ff *PostProcessingBuilder) NoEmbedSubs() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "embedsubtitles",
		Flag: "--no-embed-subs",
		Args: nil,
	})
	return ff
}

// Embed thumbnail in the video as cover art
//
// EmbedThumbnail maps to cli flags: --embed-thumbnail.
func (ff *PostProcessingBuilder) EmbedThumbnail() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "embedthumbnail",
		Flag: "--embed-thumbnail",
		Args: nil,
	})
	return ff
}

// Do not embed thumbnail (default)
//
// NoEmbedThumbnail maps to cli flags: --no-embed-thumbnail.
func (ff *PostProcessingBuilder) NoEmbedThumbnail() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "embedthumbnail",
		Flag: "--no-embed-thumbnail",
		Args: nil,
	})
	return ff
}

// Embed metadata to the video file. Also embeds chapters/infojson if present
// unless --no-embed-chapters/--no-embed-info-json are used (Alias: --add-metadata)
//
// EmbedMetadata maps to cli flags: --embed-metadata/--add-metadata.
func (ff *PostProcessingBuilder) EmbedMetadata() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "addmetadata",
		Flag: "--embed-metadata",
		Args: nil,
	})
	return ff
}

// Do not add metadata to file (default) (Alias: --no-add-metadata)
//
// NoEmbedMetadata maps to cli flags: --no-embed-metadata/--no-add-metadata.
func (ff *PostProcessingBuilder) NoEmbedMetadata() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "addmetadata",
		Flag: "--no-embed-metadata",
		Args: nil,
	})
	return ff
}

// Add chapter markers to the video file (Alias: --add-chapters)
//
// EmbedChapters maps to cli flags: --embed-chapters/--add-chapters.
func (ff *PostProcessingBuilder) EmbedChapters() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "addchapters",
		Flag: "--embed-chapters",
		Args: nil,
	})
	return ff
}

// Do not add chapter markers (default) (Alias: --no-add-chapters)
//
// NoEmbedChapters maps to cli flags: --no-embed-chapters/--no-add-chapters.
func (ff *PostProcessingBuilder) NoEmbedChapters() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "addchapters",
		Flag: "--no-embed-chapters",
		Args: nil,
	})
	return ff
}

// Embed the infojson as an attachment to mkv/mka video files
//
// EmbedInfoJson maps to cli flags: --embed-info-json.
func (ff *PostProcessingBuilder) EmbedInfoJson() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "embed_infojson",
		Flag: "--embed-info-json",
		Args: nil,
	})
	return ff
}

// Do not embed the infojson as an attachment to the video file
//
// NoEmbedInfoJson maps to cli flags: --no-embed-info-json.
func (ff *PostProcessingBuilder) NoEmbedInfoJson() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "embed_infojson",
		Flag: "--no-embed-info-json",
		Args: nil,
	})
	return ff
}

// MetadataFromTitle maps to cli flags: --metadata-from-title=FORMAT.
func (ff *PostProcessingBuilder) MetadataFromTitle(format string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "metafromtitle",
		Flag: "--metadata-from-title",
		Args: []string{format},
	})
	return ff
}

// Write metadata to the video file's xattrs (using dublin core and xdg standards)
//
// Xattrs maps to cli flags: --xattrs/--xattr.
func (ff *PostProcessingBuilder) Xattrs() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "xattrs",
		Flag: "--xattrs",
		Args: nil,
	})
	return ff
}

var concatPlaylistChoices = []string{
	"never",
	"always",
	"multi_video",
}

// Concatenate videos in a playlist. One of "never", "always", or "multi_video"
// (default; only when the videos form a single show). All the video files must
// have same codecs and number of streams to be concatable. The "pl_video:" prefix
// can be used with "--paths" and "--output" to set the output filename for the
// concatenated files. See "OUTPUT TEMPLATE" for details
//
// ConcatPlaylist maps to cli flags: --concat-playlist=POLICY.
func (ff *PostProcessingBuilder) ConcatPlaylist(policy string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "concat_playlist",
		Flag: "--concat-playlist",
		Args: []string{policy},
	})
	return ff
}

var fixupChoices = []string{
	"never",
	"ignore",
	"warn",
	"detect_or_warn",
	"force",
}

// Automatically correct known faults of the file. One of never (do nothing), warn
// (only emit a warning), detect_or_warn (the default; fix file if we can, warn
// otherwise), force (try fixing even if file already exists)
//
// Fixup maps to cli flags: --fixup=POLICY.
func (ff *PostProcessingBuilder) Fixup(policy string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "fixup",
		Flag: "--fixup",
		Args: []string{policy},
	})
	return ff
}

// PreferAvconv sets the "prefer-avconv" flag to "false".
//
// PreferAvconv maps to cli flags: --prefer-avconv/--no-prefer-ffmpeg.
func (ff *PostProcessingBuilder) PreferAvconv() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "prefer_ffmpeg",
		Flag: "--prefer-avconv",
		Args: nil,
	})
	return ff
}

// PreferFfmpeg sets the "prefer-ffmpeg" flag to "true".
//
// PreferFfmpeg maps to cli flags: --prefer-ffmpeg/--no-prefer-avconv.
func (ff *PostProcessingBuilder) PreferFfmpeg() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "prefer_ffmpeg",
		Flag: "--prefer-ffmpeg",
		Args: nil,
	})
	return ff
}

// Location of the ffmpeg binary; either the path to the binary or its containing
// directory
//
// FfmpegLocation maps to cli flags: --ffmpeg-location/--avconv-location=PATH.
func (ff *PostProcessingBuilder) FfmpegLocation(path string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "ffmpeg_location",
		Flag: "--ffmpeg-location",
		Args: []string{path},
	})
	return ff
}

// Convert the subtitles to another format (currently supported: ass, lrc, srt,
// vtt) (Alias: --convert-subtitles)
//
// ConvertSubs maps to cli flags: --convert-subs/--convert-sub/--convert-subtitles=FORMAT.
func (ff *PostProcessingBuilder) ConvertSubs(format string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "convertsubtitles",
		Flag: "--convert-subs",
		Args: []string{format},
	})
	return ff
}

// Convert the thumbnails to another format (currently supported: jpg, png, webp).
// You can specify multiple rules using similar syntax as --remux-video
//
// ConvertThumbnails maps to cli flags: --convert-thumbnails=FORMAT.
func (ff *PostProcessingBuilder) ConvertThumbnails(format string) *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "convertthumbnails",
		Flag: "--convert-thumbnails",
		Args: []string{format},
	})
	return ff
}

// Split video into multiple files based on internal chapters. The "chapter:"
// prefix can be used with "--paths" and "--output" to set the output filename for
// the split files. See "OUTPUT TEMPLATE" for details
//
// SplitChapters maps to cli flags: --split-chapters/--split-tracks.
func (ff *PostProcessingBuilder) SplitChapters() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "split_chapters",
		Flag: "--split-chapters",
		Args: nil,
	})
	return ff
}

// Do not split video based on chapters (default)
//
// NoSplitChapters maps to cli flags: --no-split-chapters/--no-split-tracks.
func (ff *PostProcessingBuilder) NoSplitChapters() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "split_chapters",
		Flag: "--no-split-chapters",
		Args: nil,
	})
	return ff
}

// Force keyframes at cuts when downloading/splitting/removing sections. This is
// slow due to needing a re-encode, but the resulting video may have fewer
// artifacts around the cuts
//
// ForceKeyframesAtCuts maps to cli flags: --force-keyframes-at-cuts.
func (ff *PostProcessingBuilder) ForceKeyframesAtCuts() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "force_keyframes_at_cuts",
		Flag: "--force-keyframes-at-cuts",
		Args: nil,
	})
	return ff
}

// Do not force keyframes around the chapters when cutting/splitting (default)
//
// NoForceKeyframesAtCuts maps to cli flags: --no-force-keyframes-at-cuts.
func (ff *PostProcessingBuilder) NoForceKeyframesAtCuts() *PostProcessingBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "force_keyframes_at_cuts",
		Flag: "--no-force-keyframes-at-cuts",
		Args: nil,
	})
	return ff
}
