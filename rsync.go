package grsync

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Rsync is wrapper under rsync
type Rsync struct {
	Source      string
	Destination string

	cmd *exec.Cmd
}

// RsyncOptions for rsync
type RsyncOptions struct {
	// Verbose increase verbosity
	Verbose bool
	// Quet suppress non-error messages
	Quiet bool
	// Checksum skip based on checksum, not mod-time & size
	Checksum bool
	// Archve is archive mode; equals -rlptgoD (no -H,-A,-X)
	Archive bool
	// Recurse into directories
	Recursive bool
	// Relative option to use relative path names
	Relative bool
	// NoImliedDirs don't send implied dirs with --relative
	NoImpliedDirs bool
	// Update skip files that are newer on the receiver
	Update bool
	// Inplace update destination files in-place
	Inplace bool
	// Append data onto shorter files
	Append bool
	// AppendVerify --append w/old data in file checksum
	AppendVerify bool
	// Dirs transfer directories without recursing
	Dirs bool
	// Links copy symlinks as symlinks
	Links bool
	// CopyLinks transform symlink into referent file/dir
	CopyLinks bool
	// CopyUnsafeLinks only "unsafe" symlinks are transformed
	CopyUnsafeLinks bool
	// SafeLinks ignore symlinks that point outside the tree
	SafeLinks bool
	// CopyDirLinks transform symlink to dir into referent dir
	CopyDirLinks bool
	// KeepDirLinks treat symlinked dir on receiver as dir
	KeepDirLinks bool
	// HardLinks preserve hard links
	HardLinks bool
	// Perms preserve permissions
	Perms bool
	// Executability preserve executability
	Executability bool
	// CHMOD affect file and/or directory permissions
	CHMOD os.FileMode
	// Acls preserve ACLs (implies -p)
	ACLs bool
	// XAttrs preserve extended attributes
	XAttrs bool
	// Owner preserve owner (super-user only)
	Owner bool
	// Group preserve group
	Group bool
	// Devices preserve device files (super-user only)
	Devices bool
	// Specials preserve special files
	Specials bool
	// Times preserve modification times
	Times bool
	// omit directories from --times
	OmitDirTimes bool
	// Super receiver attempts super-user activities
	Super bool
	// FakeSuper store/recover privileged attrs using xattrs
	FakeSuper bool
	// Sparce handle sparse files efficiently
	Sparse bool
	// DryRun perform a trial run with no changes made
	DryRun bool
	// WholeFile copy files whole (w/o delta-xfer algorithm)
	WholeFile bool
	// OneFileSystem don't cross filesystem boundaries
	OneFileSystem bool
	// BlockSize block-size=SIZE force a fixed checksum block-size
	BlockSize int
	// Rsh -rsh=COMMAND specify the remote shell to use
	Rsh string
	// RsyncProgramm rsync-path=PROGRAM specify the rsync to run on remote machine
	RsyncProgramm string
	// Existing skip creating new files on receiver
	Existing bool
	// IgnoreExisting skip updating files that exist on receiver
	IgnoreExisting bool
	// RemoveSourceFiles sender removes synchronized files (non-dir)
	RemoveSourceFiles bool
	// Delete delete extraneous files from dest dirs
	Delete bool
	// DeleteBefore receiver deletes before transfer, not during
	DeleteBefore bool
	// DeleteDuring receiver deletes during the transfer
	DeleteDuring bool
	// DeleteDelay find deletions during, delete after
	DeleteDelay bool
	// DeleteAfter receiver deletes after transfer, not during
	DeleteAfter bool
	// DeleteExcluded also delete excluded files from dest dirs
	DeleteExcluded bool
	// IgnoreErrors delete even if there are I/O errors
	IgnoreErrors bool
	// Force deletion of dirs even if not empty
	Force bool
	// MaxDelete max-delete=NUM don't delete more than NUM files
	MaxDelete int
	// MaxSize max-size=SIZE don't transfer any file larger than SIZE
	MaxSize int
	// MinSize don't transfer any file smaller than SIZE
	MinSize int
	// Partial keep partially transferred files
	Partial bool
	// PartialDir partial-dir=DIR
	PartialDir string
	// DelayUpdates put all updated files into place at end
	DelayUpdates bool
	// PruneEmptyDirs prune empty directory chains from file-list
	PruneEmptyDirs bool
	// NumericIDs don't map uid/gid values by user/group name
	NumericIDs bool
	// Timeout timeout=SECONDS set I/O timeout in seconds
	Timeout int
	// Contimeout contimeout=SECONDS set daemon connection timeout in seconds
	Contimeout int
	// IgnoreTimes don't skip files that match size and time
	IgnoreTimes bool
	// SizeOnly skip files that match in size
	SizeOnly bool
	// ModifyWindow modify-window=NUM compare mod-times with reduced accuracy
	ModifyWindow bool
	// TempDir temp-dir=DIR create temporary files in directory DIR
	TempDir string
	// Fuzzy find similar file for basis if no dest file
	Fuzzy bool
	// CompareDest compare-dest=DIR also compare received files relative to DIR
	CompareDest string
	// CopyDest copy-dest=DIR ... and include copies of unchanged files
	CopyDest string
	// LinkDest link-dest=DIR hardlink to files in DIR when unchanged
	LinkDest string
	// Compress file data during the transfer
	Compress bool
	// CompressLevel explicitly set compression level
	CompressLevel int
	// SkipCompress skip-compress=LIST skip compressing files with suffix in LIST
	SkipCompress []string
	// CVSExclude auto-ignore files in the same way CVS does
	CVSExclude bool
	// Stats give some file-transfer stats
	Stats bool
	// HumanReadable output numbers in a human-readable format
	HumanReadable bool
	// Progress show progress during transfer
	Progress bool
	// Info
	Info string
	// Exclude --exclude="", exclude remote paths.
	Exclude []string
	// Include --include="", include remote paths.
	Include []string
	// Filter --filter="", include filter rule.
	Filter string
	// Chown --chown="", chown on receipt.
	Chown string

	// --no-OPTION flags.
	No *RsyncOptions

	// ipv4
	IPv4 bool
	// ipv6
	IPv6 bool

	//out-format
	OutFormat bool
}

// StdoutPipe returns a pipe that will be connected to the command's
// standard output when the command starts.
func (r Rsync) StdoutPipe() (io.ReadCloser, error) {
	return r.cmd.StdoutPipe()
}

// StderrPipe returns a pipe that will be connected to the command's
// standard error when the command starts.
func (r Rsync) StderrPipe() (io.ReadCloser, error) {
	return r.cmd.StderrPipe()
}

// Run start rsync task
func (r Rsync) Run() error {
	if !isExist(r.Destination) {
		if err := createDir(r.Destination); err != nil {
			return err
		}
	}

	if err := r.cmd.Start(); err != nil {
		return err
	}

	return r.cmd.Wait()
}

// NewRsync returns task with described options
func NewRsync(source, destination string, options RsyncOptions) *Rsync {
	arguments := append(GetArguments(options), source, destination)
	return &Rsync{
		Source:      source,
		Destination: destination,
		cmd:         exec.Command("rsync", arguments...),
	}
}

func GetArguments(options RsyncOptions) []string {
	args := GetArgsPrefix(options, "--")
	if options.No != nil {
		args = append(args, GetArgsPrefix(*options.No, "--no-")...)
	}
	return args
}

func GetArgsPrefix(options RsyncOptions, prefix string) []string {
	arguments := []string{}
	if options.Verbose {
		arguments = append(arguments, fmt.Sprintf("%sverbose", prefix))
	}

	if options.Checksum {
		arguments = append(arguments, fmt.Sprintf("%schecksum", prefix))
	}

	if options.Quiet {
		arguments = append(arguments, fmt.Sprintf("%squiet", prefix))
	}

	if options.Archive {
		arguments = append(arguments, fmt.Sprintf("%sarchive", prefix))
	}

	if options.Recursive {
		arguments = append(arguments, fmt.Sprintf("%srecursive", prefix))
	}

	if options.Relative {
		arguments = append(arguments, fmt.Sprintf("%srelative", prefix))
	}

	if options.NoImpliedDirs {
		arguments = append(arguments, fmt.Sprintf("%sno-implied-dirs", prefix))
	}

	if options.Update {
		arguments = append(arguments, fmt.Sprintf("%supdate", prefix))
	}

	if options.Inplace {
		arguments = append(arguments, fmt.Sprintf("%sinplace", prefix))
	}

	if options.Append {
		arguments = append(arguments, fmt.Sprintf("%sappend", prefix))
	}

	if options.AppendVerify {
		arguments = append(arguments, fmt.Sprintf("%sappend-verify", prefix))
	}

	if options.Dirs {
		arguments = append(arguments, fmt.Sprintf("%sdirs", prefix))
	}

	if options.Links {
		arguments = append(arguments, fmt.Sprintf("%slinks", prefix))
	}

	if options.CopyLinks {
		arguments = append(arguments, fmt.Sprintf("%scopy-links", prefix))
	}

	if options.CopyUnsafeLinks {
		arguments = append(arguments, fmt.Sprintf("%scopy-unsafe-links", prefix))
	}

	if options.SafeLinks {
		arguments = append(arguments, fmt.Sprintf("%ssafe-links", prefix))
	}

	if options.CopyDirLinks {
		arguments = append(arguments, fmt.Sprintf("%scopy-dir-links", prefix))
	}

	if options.KeepDirLinks {
		arguments = append(arguments, fmt.Sprintf("%skeep-dir-links", prefix))
	}

	if options.HardLinks {
		arguments = append(arguments, fmt.Sprintf("%shard-links", prefix))
	}

	if options.Perms {
		arguments = append(arguments, fmt.Sprintf("%sperms", prefix))
	}

	if options.Executability {
		arguments = append(arguments, fmt.Sprintf("%sexecutability", prefix))
	}

	if options.ACLs {
		arguments = append(arguments, fmt.Sprintf("%sacls", prefix))
	}

	if options.XAttrs {
		arguments = append(arguments, fmt.Sprintf("%sxattrs", prefix))
	}

	if options.Owner {
		arguments = append(arguments, fmt.Sprintf("%sowner", prefix))
	}

	if options.Group {
		arguments = append(arguments, fmt.Sprintf("%sgroup", prefix))
	}

	if options.Devices {
		arguments = append(arguments, fmt.Sprintf("%sdevices", prefix))
	}

	if options.Specials {
		arguments = append(arguments, fmt.Sprintf("%sspecials", prefix))
	}

	if options.Times {
		arguments = append(arguments, fmt.Sprintf("%stimes", prefix))
	}

	if options.OmitDirTimes {
		arguments = append(arguments, fmt.Sprintf("%somit-dir-times", prefix))
	}

	if options.Super {
		arguments = append(arguments, fmt.Sprintf("%ssuper", prefix))
	}

	if options.FakeSuper {
		arguments = append(arguments, fmt.Sprintf("%sfake-super", prefix))
	}

	if options.Sparse {
		arguments = append(arguments, fmt.Sprintf("%ssparse", prefix))
	}

	if options.DryRun {
		arguments = append(arguments, fmt.Sprintf("%sdry-run", prefix))
	}

	if options.WholeFile {
		arguments = append(arguments, fmt.Sprintf("%swhole-file", prefix))
	}

	if options.OneFileSystem {
		arguments = append(arguments, fmt.Sprintf("%sone-file-system", prefix))
	}

	if options.BlockSize > 0 {
		arguments = append(arguments, fmt.Sprintf("%sblock-size", prefix), strconv.Itoa(options.BlockSize))
	}

	if options.Rsh != "" {
		arguments = append(arguments, fmt.Sprintf("%srsh", prefix), options.Rsh)
	}

	if options.RsyncProgramm != "" {
		arguments = append(arguments, fmt.Sprintf("%srsync-programm", prefix), options.RsyncProgramm)
	}

	if options.Existing {
		arguments = append(arguments, fmt.Sprintf("%sexisting", prefix))
	}

	if options.IgnoreExisting {
		arguments = append(arguments, fmt.Sprintf("%signore-existing", prefix))
	}

	if options.RemoveSourceFiles {
		arguments = append(arguments, fmt.Sprintf("%sremove-source-files", prefix))
	}

	if options.Delete {
		arguments = append(arguments, fmt.Sprintf("%sdelete", prefix))
	}

	if options.DeleteBefore {
		arguments = append(arguments, fmt.Sprintf("%sdelete-before", prefix))
	}

	if options.DeleteDuring {
		arguments = append(arguments, fmt.Sprintf("%sdelete-during", prefix))
	}

	if options.DeleteDelay {
		arguments = append(arguments, fmt.Sprintf("%sdelete-delay", prefix))
	}

	if options.DeleteAfter {
		arguments = append(arguments, fmt.Sprintf("%sdelete-after", prefix))
	}

	if options.DeleteExcluded {
		arguments = append(arguments, fmt.Sprintf("%sdelete-excluded", prefix))
	}

	if options.IgnoreErrors {
		arguments = append(arguments, fmt.Sprintf("%signore-errors", prefix))
	}

	if options.Force {
		arguments = append(arguments, fmt.Sprintf("%sforce", prefix))
	}

	if options.MaxDelete > 0 {
		arguments = append(arguments, fmt.Sprintf("%smax-delete", prefix), strconv.Itoa(options.MaxDelete))
	}

	if options.MaxSize > 0 {
		arguments = append(arguments, fmt.Sprintf("%smax-size", prefix), strconv.Itoa(options.MaxSize))
	}

	if options.MinSize > 0 {
		arguments = append(arguments, fmt.Sprintf("%smin-size", prefix), strconv.Itoa(options.MinSize))
	}

	if options.Partial {
		arguments = append(arguments, fmt.Sprintf("%spartial", prefix))
	}

	if options.PartialDir != "" {
		arguments = append(arguments, fmt.Sprintf("%spartial-dir", prefix), options.PartialDir)
	}

	if options.DelayUpdates {
		arguments = append(arguments, fmt.Sprintf("%sdelay-updates", prefix))
	}

	if options.PruneEmptyDirs {
		arguments = append(arguments, fmt.Sprintf("%sprune-empty-dirs", prefix))
	}

	if options.NumericIDs {
		arguments = append(arguments, fmt.Sprintf("%snumeric-ids", prefix))
	}

	if options.Timeout > 0 {
		arguments = append(arguments, fmt.Sprintf("%stimeout", prefix), strconv.Itoa(options.Timeout))
	}

	if options.Contimeout > 0 {
		arguments = append(arguments, fmt.Sprintf("%scontimeout", prefix), strconv.Itoa(options.Contimeout))
	}

	if options.IgnoreTimes {
		arguments = append(arguments, fmt.Sprintf("%signore-times", prefix))
	}

	if options.SizeOnly {
		arguments = append(arguments, fmt.Sprintf("%ssize-only", prefix))
	}

	if options.ModifyWindow {
		arguments = append(arguments, fmt.Sprintf("%smodify-window", prefix))
	}

	if options.TempDir != "" {
		arguments = append(arguments, fmt.Sprintf("%stemp-dir", prefix), options.TempDir)
	}

	if options.Fuzzy {
		arguments = append(arguments, fmt.Sprintf("%sfuzzy", prefix))
	}

	if options.CompareDest != "" {
		arguments = append(arguments, fmt.Sprintf("%scompare-dest", prefix), options.CompareDest)
	}

	if options.CopyDest != "" {
		arguments = append(arguments, fmt.Sprintf("%scopy-dest", prefix), options.CopyDest)
	}

	if options.LinkDest != "" {
		arguments = append(arguments, fmt.Sprintf("%slink-dest", prefix), options.LinkDest)
	}

	if options.Compress {
		arguments = append(arguments, fmt.Sprintf("%scompress", prefix))
	}

	if options.CompressLevel > 0 {
		arguments = append(arguments, fmt.Sprintf("%scompress-level", prefix), strconv.Itoa(options.CompressLevel))
	}

	if len(options.SkipCompress) > 0 {
		arguments = append(arguments, fmt.Sprintf("%sskip-compress", prefix), strings.Join(options.SkipCompress, ","))
	}

	if options.CVSExclude {
		arguments = append(arguments, fmt.Sprintf("%scvs-exclude", prefix))
	}

	if options.Stats {
		arguments = append(arguments, fmt.Sprintf("%sstats", prefix))
	}

	if options.HumanReadable {
		arguments = append(arguments, fmt.Sprintf("%shuman-readable", prefix))
	}

	if options.Progress {
		arguments = append(arguments, fmt.Sprintf("%sprogress", prefix))
	}

	if options.IPv4 {
		arguments = append(arguments, fmt.Sprintf("%sipv4", prefix))
	}

	if options.IPv6 {
		arguments = append(arguments, fmt.Sprintf("%sipv6", prefix))
	}

	if options.Info != "" {
		arguments = append(arguments, fmt.Sprintf("%sinfo", prefix), options.Info)
	}

	if options.OutFormat {
		arguments = append(arguments, fmt.Sprintf("%sout-format=\"%%n\"", prefix))
	}

	if len(options.Exclude) > 0 {
		for _, pattern := range options.Exclude {
			arguments = append(arguments, fmt.Sprintf("%sexclude=%s", prefix, pattern))
		}
	}

	if len(options.Include) > 0 {
		for _, pattern := range options.Include {
			arguments = append(arguments, fmt.Sprintf("%sinclude=%s", prefix, pattern))
		}
	}

	if options.Filter != "" {
		arguments = append(arguments, fmt.Sprintf("%sfilter=%s", prefix, options.Filter))
	}

	if options.Chown != "" {
		arguments = append(arguments, fmt.Sprintf("%schown=%s", prefix, options.Chown))
	}

	return arguments
}

func createDir(dir string) error {
	cmd := exec.Command("mkdir", "-p", dir)
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func isExist(p string) bool {
	stat, err := os.Stat(p)
	return os.IsExist(err) && stat.IsDir()
}
