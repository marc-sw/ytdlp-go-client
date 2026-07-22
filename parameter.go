package ytdlp

import "fmt"

// Parameter represents a single yt-dlp CLI option with an optional value.
type Parameter struct {
	flag  Flag
	value string
}

func newParameter(flag Flag, value string) Parameter {
	return Parameter{flag: flag, value: value}
}

func (p Parameter) apply(c *Args) {
	c.args = append(c.args, p.flag.String(), p.value)
}

// When is used for options that accept a when/condition style value.
type When string

const (
	WhenAfterExtraction When = "pre_process"
	WhenAfterFilter     When = "after_filter"
	WhenAfterFormat     When = "video"
	WhenBeforeDownload  When = "before_dl"
	WhenAfterDownload   When = "post_process"
	WhenAfterMove       When = "after_move"
	WhenAfterVideo      When = "after_video"
	WhenAfterPlaylist   When = "playlist"
)

func (w When) String() string { return string(w) }

// Audio represents supported audio formats for --audio-format.
type Audio string

const (
	AudioAAC    Audio = "aac"
	AudioALAC   Audio = "alac"
	AudioFLAC   Audio = "flac"
	AudioM4A    Audio = "m4a"
	AudioMP3    Audio = "mp3"
	AudioOPUS   Audio = "opus"
	AudioVORBIS Audio = "vorbis"
	AudioWAV    Audio = "wav"
)

func (f Audio) String() string { return string(f) }

// Video represents supported video formats for remux/recode options.
type Video string

const (
	VideoAVI    Video = "avi"
	VideoFLV    Video = "flv"
	VideoGIF    Video = "gif"
	VideoMKV    Video = "mkv"
	VideoMOV    Video = "mov"
	VideoMP4    Video = "mp4"
	VideoWEBM   Video = "webm"
	VideoAAC    Video = "aac"
	VideoAIFF   Video = "aiff"
	VideoALAC   Video = "alac"
	VideoFLAC   Video = "flac"
	VideoM4A    Video = "m4a"
	VideoMKA    Video = "mka"
	VideoMP3    Video = "mp3"
	VideoOGG    Video = "ogg"
	VideoOPUS   Video = "opus"
	VideoVORBIS Video = "vorbis"
	VideoWAV    Video = "wav"
)

func (f Video) String() string { return string(f) }

// Preset represents preset aliases.
type Preset string

const (
	PresetMP3   Preset = "mp3"
	PresetAAC   Preset = "aac"
	PresetMP4   Preset = "mp4"
	PresetMKV   Preset = "mkv"
	PresetSleep Preset = "sleep"
)

func (p Preset) String() string { return string(p) }

// BiParameter represents a two-value option such as --print-to-file.
type BiParameter struct {
	flag   Flag
	value1 string
	value2 string
}

func newBiParameter(flag Flag, value1, value2 string) BiParameter {
	return BiParameter{flag: flag, value1: value1, value2: value2}
}

func (p BiParameter) apply(c *Args) {
	c.args = append(c.args, p.flag.String(), p.value1, p.value2)
}

// General
func UpdateTo(channel, tag string) Parameter {
	return newParameter(FLAG_UPDATE_TO, fmt.Sprintf("%s@%s", channel, tag))
}

func UseExtractors(names string) Parameter {
	return newParameter(FLAG_USE_EXTRACTORS, names)
}

func DefaultSearch(prefix string) Parameter {
	return newParameter(FLAG_DEFAULT_SEARCH, prefix)
}

func ConfigLocations(path string) Parameter {
	return newParameter(FLAG_CONFIG_LOCATIONS, path)
}

func PluginDirs(path string) Parameter {
	return newParameter(FLAG_PLUGIN_DIRS, path)
}

func WaitForVideo(min int) Parameter {
	return newParameter(FLAG_WAIT_FOR_VIDEO, fmt.Sprintf("%d", min))
}

func WaitForVideoRange(min, max int) Parameter {
	return newParameter(FLAG_WAIT_FOR_VIDEO, fmt.Sprintf("%d-%d", min, max))
}

func Color(stream, policy string) Parameter {
	return newParameter(FLAG_COLOR, fmt.Sprintf("%s:%s", stream, policy))
}

func ColorPolicy(policy string) Parameter {
	return newParameter(FLAG_COLOR, policy)
}

func CompatOptions(opts string) Parameter {
	return newParameter(FLAG_COMPAT_OPTIONS, opts)
}

// Network
func Proxy(url string) Parameter {
	return newParameter(FLAG_PROXY, url)
}

func SocketTimeout(seconds int) Parameter {
	return newParameter(FLAG_SOCKET_TIMEOUT, fmt.Sprintf("%d", seconds))
}

func SourceAddress(ip string) Parameter {
	return newParameter(FLAG_SOURCE_ADDRESS, ip)
}

func Impersonate(client, os string) Parameter {
	return newParameter(FLAG_IMPERSONATE, fmt.Sprintf("%s:%s", client, os))
}

func ImpersonateClient(client string) Parameter {
	return newParameter(FLAG_IMPERSONATE, client)
}

// Geo-restriction
func GeoVerificationProxy(url string) Parameter {
	return newParameter(FLAG_GEO_VERIFICATION_PROXY, url)
}

func FakeForward(header string) Parameter {
	return newParameter(FLAG_XFF, header)
}

// Video Selection
func PlaylistItems(items string) Parameter {
	return newParameter(FLAG_PLAYLIST_ITEMS, items)
}

func PlaylistItemsRange(start, end int) Parameter {
	return newParameter(FLAG_PLAYLIST_ITEMS, fmt.Sprintf("%d:%d", start, end))
}

func MinFilesize(filesize string) Parameter {
	return newParameter(FLAG_MIN_FILESIZE, filesize)
}

func MaxFilesize(filesize string) Parameter {
	return newParameter(FLAG_MAX_FILESIZE, filesize)
}

func Date(date string) Parameter {
	return newParameter(FLAG_DATE, date)
}

func DateBefore(date string) Parameter {
	return newParameter(FLAG_DATEBEFORE, date)
}

func DateAfter(date string) Parameter {
	return newParameter(FLAG_DATEAFTER, date)
}

func MatchFilters(filter string) Parameter {
	return newParameter(FLAG_MATCH_FILTERS, filter)
}

func BreakMatchFilters(filter string) Parameter {
	return newParameter(FLAG_BREAK_MATCH_FILTERS, filter)
}

func AgeLimit(years int) Parameter {
	return newParameter(FLAG_AGE_LIMIT, fmt.Sprintf("%d", years))
}

func DownloadArchive(file string) Parameter {
	return newParameter(FLAG_DOWNLOAD_ARCHIVE, file)
}

func MaxDownloads(number int) Parameter {
	return newParameter(FLAG_MAX_DOWNLOADS, fmt.Sprintf("%d", number))
}

func SkipPlaylistAfterErrors(errors int) Parameter {
	return newParameter(FLAG_SKIP_PLAYLIST_AFTER_ERRORS, fmt.Sprintf("%d", errors))
}

// Download Options
func ConcurrentFragments(fragments int) Parameter {
	return newParameter(FLAG_CONCURRENT_FRAGMENTS, fmt.Sprintf("%d", fragments))
}

func LimitRate(rate string) Parameter {
	return newParameter(FLAG_LIMIT_RATE, rate)
}

func ThrottledRate(rate string) Parameter {
	return newParameter(FLAG_THROTTLED_RATE, rate)
}

func Retries(retries int) Parameter {
	return newParameter(FLAG_RETRIES, fmt.Sprintf("%d", retries))
}

func FileAccessRetries(retries int) Parameter {
	return newParameter(FLAG_FILE_ACCESS_RETRIES, fmt.Sprintf("%d", retries))
}

func FragmentRetries(retries int) Parameter {
	return newParameter(FLAG_FRAGMENT_RETRIES, fmt.Sprintf("%d", retries))
}

func RetrySleep(typeValue, expression string) Parameter {
	return newParameter(FLAG_RETRY_SLEEP, fmt.Sprintf("%s:%s", typeValue, expression))
}

func RetrySleepExpression(expression string) Parameter {
	return newParameter(FLAG_RETRY_SLEEP, expression)
}

func BufferSize(size string) Parameter {
	return newParameter(FLAG_BUFFER_SIZE, size)
}

func HttpChunkSize(size string) Parameter {
	return newParameter(FLAG_HTTP_CHUNK_SIZE, size)
}

func DownloadSections(regex string) Parameter {
	return newParameter(FLAG_DOWNLOAD_SECTIONS, regex)
}

func Downloader(proto, name string) Parameter {
	return newParameter(FLAG_DOWNLOADER, fmt.Sprintf("%s:%s", proto, name))
}

func DownloaderName(name string) Parameter {
	return newParameter(FLAG_DOWNLOADER, name)
}

func DownloaderArgs(name, args string) Parameter {
	return newParameter(FLAG_DOWNLOADER_ARGS, fmt.Sprintf("%s:%s", name, args))
}

func BatchFile(file string) Parameter {
	return newParameter(FLAG_BATCH_FILE, file)
}

func Paths(types, path string) Parameter {
	return newParameter(FLAG_PATHS, fmt.Sprintf("%s:%s", types, path))
}

func PathsSingle(path string) Parameter {
	return newParameter(FLAG_PATHS, path)
}

func Output(types, template string) Parameter {
	return newParameter(FLAG_OUTPUT, fmt.Sprintf("%s:%s", types, template))
}

func OutputTemplate(template string) Parameter {
	return newParameter(FLAG_OUTPUT, template)
}

func OutputNaPlaceholder(text string) Parameter {
	return newParameter(FLAG_OUTPUT_NA_PLACEHOLDER, text)
}

func TrimFilenames(length int) Parameter {
	return newParameter(FLAG_TRIM_FILENAMES, fmt.Sprintf("%d", length))
}

func LoadInfoJson(file string) Parameter {
	return newParameter(FLAG_LOAD_INFO_JSON, file)
}

func Cookies(file string) Parameter {
	return newParameter(FLAG_COOKIES, file)
}

func CookiesFromBrowser(browser string) Parameter {
	return newParameter(FLAG_COOKIES_FROM_BROWSER, browser)
}

func CacheDir(directory string) Parameter {
	return newParameter(FLAG_CACHE_DIR, directory)
}

// Verbosity and Simulation Options
func Print(when When, template string) Parameter {
	return newParameter(FLAG_PRINT, fmt.Sprintf("%s:%s", when, template))
}

func PrintTemplate(template string) Parameter {
	return newParameter(FLAG_PRINT, template)
}

func PrintToFile(when When, template, file string) BiParameter {
	return newBiParameter(FLAG_PRINT_TO_FILE, fmt.Sprintf("%s:%s", when, template), file)
}

func PrintToFileTemplate(template, file string) BiParameter {
	return newBiParameter(FLAG_PRINT_TO_FILE, template, file)
}

func ProgressTemplate(types, template string) Parameter {
	return newParameter(FLAG_PROGRESS_TEMPLATE, fmt.Sprintf("%s:%s", types, template))
}

func ProgressTemplateValue(template string) Parameter {
	return newParameter(FLAG_PROGRESS_TEMPLATE, template)
}

func ProgressDelta(seconds float32) Parameter {
	return newParameter(FLAG_PROGRESS_DELTA, fmt.Sprintf("%g", seconds))
}

// Workarounds
func Encoding(encoding string) Parameter {
	return newParameter(FLAG_ENCODING, encoding)
}

func AddHeaders(header, value string) Parameter {
	return newParameter(FLAG_ADD_HEADERS, fmt.Sprintf("%s:%s", header, value))
}

func SleepRequests(seconds float32) Parameter {
	return newParameter(FLAG_SLEEP_REQUESTS, fmt.Sprintf("%g", seconds))
}

func SleepInterval(seconds float32) Parameter {
	return newParameter(FLAG_SLEEP_INTERVAL, fmt.Sprintf("%g", seconds))
}

func SleepMaxInterval(seconds float32) Parameter {
	return newParameter(FLAG_MAX_SLEEP_INTERVAL, fmt.Sprintf("%g", seconds))
}

func SleepSubtitles(seconds float32) Parameter {
	return newParameter(FLAG_SLEEP_SUBTITLES, fmt.Sprintf("%g", seconds))
}

// Video Format Options
func Format(format string) Parameter {
	return newParameter(FLAG_FORMAT, format)
}

func FormatSort(sortOrder string) Parameter {
	return newParameter(FLAG_FORMAT_SORT, sortOrder)
}

func MergeOutputFormat(format string) Parameter {
	return newParameter(FLAG_MERGE_OUTPUT_FORMAT, format)
}

// Subtitle Options
func SubFormat(format string) Parameter {
	return newParameter(FLAG_SUB_FORMAT, format)
}

func SubLangs(langs string) Parameter {
	return newParameter(FLAG_SUB_LANGS, langs)
}

// Authentication Options
func Username(username string) Parameter {
	return newParameter(FLAG_USERNAME, username)
}

func Password(password string) Parameter {
	return newParameter(FLAG_PASSWORD, password)
}

func TwoFactor(code string) Parameter {
	return newParameter(FLAG_TWO_FACTOR, code)
}

func NetrcLocation(path string) Parameter {
	return newParameter(FLAG_NETRC_LOCATION, path)
}

func NetrcCmd(cmd string) Parameter {
	return newParameter(FLAG_NETRC_CMD, cmd)
}

func VideoPassword(password string) Parameter {
	return newParameter(FLAG_VIDEO_PASSWORD, password)
}

func ApMso(mso string) Parameter {
	return newParameter(FLAG_AP_MSO, mso)
}

func ApUsername(username string) Parameter {
	return newParameter(FLAG_AP_USERNAME, username)
}

func ApPassword(password string) Parameter {
	return newParameter(FLAG_AP_PASSWORD, password)
}

func ClientCertificate(certFile string) Parameter {
	return newParameter(FLAG_CLIENT_CERTIFICATE, certFile)
}

func ClientCertificateKey(keyFile string) Parameter {
	return newParameter(FLAG_CLIENT_CERTIFICATE_KEY, keyFile)
}

func ClientCertificatePassword(password string) Parameter {
	return newParameter(FLAG_CLIENT_CERTIFICATE_PASSWORD, password)
}

// Post-Processing Options
func AudioFormat(format Audio) Parameter {
	return newParameter(FLAG_AUDIO_FORMAT, fmt.Sprintf("%s", format))
}

func AudioQuality(quality string) Parameter {
	return newParameter(FLAG_AUDIO_QUALITY, quality)
}

func RemuxVideo(format Video) Parameter {
	return newParameter(FLAG_REMUX_VIDEO, fmt.Sprintf("%s", format))
}

func RecodeVideo(format Video) Parameter {
	return newParameter(FLAG_RECODE_VIDEO, fmt.Sprintf("%s", format))
}

func PostprocessorArgs(name, args string) Parameter {
	return newParameter(FLAG_POSTPROCESSOR_ARGS, fmt.Sprintf("%s:%s", name, args))
}

func ParseMetadata(when When, from, to string) Parameter {
	return newParameter(FLAG_PARSE_METADATA, fmt.Sprintf("%s:%s:%s", when, from, to))
}

func ParseMetadataPair(from, to string) Parameter {
	return newParameter(FLAG_PARSE_METADATA, fmt.Sprintf("%s:%s", from, to))
}

func ReplaceInMetadata(when When, fields, regex, replace string) Parameter {
	return newParameter(FLAG_REPLACE_IN_METADATA, fmt.Sprintf("%s:%s %s %s", when, fields, regex, replace))
}

func ReplaceInMetadataValues(fields, regex, replace string) Parameter {
	return newParameter(FLAG_REPLACE_IN_METADATA, fmt.Sprintf("%s %s %s", fields, regex, replace))
}

func ConcatPlaylist(policy string) Parameter {
	return newParameter(FLAG_CONCAT_PLAYLIST, policy)
}

func Fixup(policy string) Parameter {
	return newParameter(FLAG_FIXUP, policy)
}

func FfmpegLocation(path string) Parameter {
	return newParameter(FLAG_FFMPEG_LOCATION, path)
}

func Exec(when When, cmd string) Parameter {
	return newParameter(FLAG_EXEC, fmt.Sprintf("%s:%s", when, cmd))
}

func ExecCommand(cmd string) Parameter {
	return newParameter(FLAG_EXEC, cmd)
}

func ConvertSubs(format string) Parameter {
	return newParameter(FLAG_CONVERT_SUBS, format)
}

func ConvertThumbnails(format string) Parameter {
	return newParameter(FLAG_CONVERT_THUMBNAILS, format)
}

func RemoveChapters(regex string) Parameter {
	return newParameter(FLAG_REMOVE_CHAPTERS, regex)
}

func UsePostprocessor(name, args string) Parameter {
	return newParameter(FLAG_USE_POSTPROCESSOR, fmt.Sprintf("%s:%s", name, args))
}

func UsePostprocessorName(name string) Parameter {
	return newParameter(FLAG_USE_POSTPROCESSOR, name)
}

// SponsorBlock Options
func SponsorblockMark(cats string) Parameter {
	return newParameter(FLAG_SPONSORBLOCK_MARK, cats)
}

func SponsorblockRemove(cats string) Parameter {
	return newParameter(FLAG_SPONSORBLOCK_REMOVE, cats)
}

func SponsorblockChapterTitle(template string) Parameter {
	return newParameter(FLAG_SPONSORBLOCK_CHAPTER_TITLE, template)
}

func SponsorblockApi(url string) Parameter {
	return newParameter(FLAG_SPONSORBLOCK_API, url)
}

// Extractor Options
func ExtractorRetries(retries string) Parameter {
	return newParameter(FLAG_EXTRACTOR_RETRIES, retries)
}

func ExtractorArgs(ieKey, args string) Parameter {
	return newParameter(FLAG_EXTRACTOR_ARGS, fmt.Sprintf("%s:%s", ieKey, args))
}

// Preset Aliases
func PresetAlias(preset Preset) Parameter {
	return newParameter(FLAG_PRESET_ALIAS, fmt.Sprintf("%s", preset))
}
