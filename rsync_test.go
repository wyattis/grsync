package grsync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArguments(t *testing.T) {
	t.Run("--verbose", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Verbose: true,
		})
		assert.Contains(t, args, "--verbose")
	})

	t.Run("--checksum", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Checksum: true,
		})
		assert.Contains(t, args, "--checksum")
	})

	t.Run("--quiet", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Quiet: true,
		})
		assert.Contains(t, args, "--quiet")
	})

	t.Run("--archive", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Archive: true,
		})
		assert.Contains(t, args, "--archive")
	})

	t.Run("--recursive", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Recursive: true,
		})
		assert.Contains(t, args, "--recursive")
	})

	t.Run("--relative", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Relative: true,
		})
		assert.Contains(t, args, "--relative")
	})

	t.Run("--no-implied-dirs", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			NoImpliedDirs: true,
		})
		assert.Contains(t, args, "--no-implied-dirs")
	})

	t.Run("--update", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Update: true,
		})
		assert.Contains(t, args, "--update")
	})

	t.Run("--inplace", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Inplace: true,
		})
		assert.Contains(t, args, "--inplace")
	})

	t.Run("--append", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Append: true,
		})
		assert.Contains(t, args, "--append")
	})

	t.Run("--append-verify", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			AppendVerify: true,
		})
		assert.Contains(t, args, "--append-verify")
	})

	t.Run("--dirs", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Dirs: true,
		})
		assert.Contains(t, args, "--dirs")
	})

	t.Run("--links", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Links: true,
		})
		assert.Contains(t, args, "--links")
	})

	t.Run("--copy-links", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			CopyLinks: true,
		})
		assert.Contains(t, args, "--copy-links")
	})

	t.Run("--copy-unsafe-links", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			CopyUnsafeLinks: true,
		})
		assert.Contains(t, args, "--copy-unsafe-links")
	})

	t.Run("--safe-links", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			SafeLinks: true,
		})
		assert.Contains(t, args, "--safe-links")
	})

	t.Run("--copy-dir-links", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			CopyDirLinks: true,
		})
		assert.Contains(t, args, "--copy-dir-links")
	})

	t.Run("--keep-dir-links", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			KeepDirLinks: true,
		})
		assert.Contains(t, args, "--keep-dir-links")
	})

	t.Run("--hard-links", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			HardLinks: true,
		})
		assert.Contains(t, args, "--hard-links")
	})

	t.Run("--perms", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Perms: true,
		})
		assert.Contains(t, args, "--perms")
	})

	t.Run("--executability", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Executability: true,
		})
		assert.Contains(t, args, "--executability")
	})

	t.Run("--acls", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			ACLs: true,
		})
		assert.Contains(t, args, "--acls")
	})

	t.Run("--xattrs", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			XAttrs: true,
		})
		assert.Contains(t, args, "--xattrs")
	})

	t.Run("--owner", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Owner: true,
		})
		assert.Contains(t, args, "--owner")
	})

	t.Run("--group", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Group: true,
		})
		assert.Contains(t, args, "--group")
	})

	t.Run("--devices", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Devices: true,
		})
		assert.Contains(t, args, "--devices")
	})

	t.Run("--specials", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Specials: true,
		})
		assert.Contains(t, args, "--specials")
	})

	t.Run("--times", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Times: true,
		})
		assert.Contains(t, args, "--times")
	})

	t.Run("--omit-dir-times", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			OmitDirTimes: true,
		})
		assert.Contains(t, args, "--omit-dir-times")
	})

	t.Run("--super", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Super: true,
		})
		assert.Contains(t, args, "--super")
	})

	t.Run("--fake-super", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			FakeSuper: true,
		})
		assert.Contains(t, args, "--fake-super")
	})

	t.Run("--sparse", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Sparse: true,
		})
		assert.Contains(t, args, "--sparse")
	})

	t.Run("--dry-run", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			DryRun: true,
		})
		assert.Contains(t, args, "--dry-run")
	})

	t.Run("--whole-file", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			WholeFile: true,
		})
		assert.Contains(t, args, "--whole-file")
	})

	t.Run("--one-file-system", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			OneFileSystem: true,
		})
		assert.Contains(t, args, "--one-file-system")
	})

	t.Run("--block-size", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			BlockSize: 1,
		})
		assert.ElementsMatch(t, args, []string{"--block-size", "1"})
	})

	t.Run("--rsh", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Rsh: "test",
		})
		assert.Contains(t, args, "--rsh", "test")
	})

	t.Run("--rsync-programm", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			RsyncProgramm: "test",
		})
		assert.Contains(t, args, "--rsync-programm", "test")
	})

	t.Run("--existing", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Existing: true,
		})
		assert.Contains(t, args, "--existing")
	})

	t.Run("--ignore-existing", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			IgnoreExisting: true,
		})
		assert.Contains(t, args, "--ignore-existing")
	})

	t.Run("--remove-source-files", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			RemoveSourceFiles: true,
		})
		assert.Contains(t, args, "--remove-source-files")
	})

	t.Run("", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Delete: true,
		})
		assert.Contains(t, args, "--delete")
	})

	t.Run("--delete", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			DeleteBefore: true,
		})
		assert.Contains(t, args, "--delete-before")
	})

	t.Run("--delete-during", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			DeleteDuring: true,
		})
		assert.Contains(t, args, "--delete-during")
	})

	t.Run("--delete-delay", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			DeleteDelay: true,
		})
		assert.Contains(t, args, "--delete-delay")
	})

	t.Run("--delete-after", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			DeleteAfter: true,
		})
		assert.Contains(t, args, "--delete-after")
	})

	t.Run("--delete-excluded", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			DeleteExcluded: true,
		})
		assert.Contains(t, args, "--delete-excluded")
	})

	t.Run("--ignore-errors", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			IgnoreErrors: true,
		})
		assert.Contains(t, args, "--ignore-errors")
	})

	t.Run("--force", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Force: true,
		})
		assert.Contains(t, args, "--force")
	})

	t.Run("--max-delete", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			MaxDelete: 1,
		})
		assert.ElementsMatch(t, args, []string{"--max-delete", "1"})
	})

	t.Run("--max-size", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			MaxSize: 1,
		})
		assert.ElementsMatch(t, args, []string{"--max-size", "1"})
	})

	t.Run("--min-size", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			MinSize: 1,
		})
		assert.ElementsMatch(t, args, []string{"--min-size", "1"})
	})

	t.Run("--partial", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Partial: true,
		})
		assert.Contains(t, args, "--partial")
	})

	t.Run("--partial-dir", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			PartialDir: "test",
		})
		assert.ElementsMatch(t, args, []string{"--partial-dir", "test"})
	})

	t.Run("--delay-updates", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			DelayUpdates: true,
		})
		assert.Contains(t, args, "--delay-updates")
	})

	t.Run("--prune-empty-dirs", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			PruneEmptyDirs: true,
		})
		assert.Contains(t, args, "--prune-empty-dirs")
	})

	t.Run("--numeric-ids", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			NumericIDs: true,
		})
		assert.Contains(t, args, "--numeric-ids")
	})

	t.Run("--timeout", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Timeout: 100,
		})
		assert.ElementsMatch(t, args, []string{"--timeout", "100"})
	})

	t.Run("--timeout", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Contimeout: 100,
		})
		assert.ElementsMatch(t, args, []string{"--contimeout", "100"})
	})

	t.Run("--ignore-times", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			IgnoreTimes: true,
		})
		assert.Contains(t, args, "--ignore-times")
	})

	t.Run("--size-only", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			SizeOnly: true,
		})
		assert.Contains(t, args, "--size-only")
	})

	t.Run("--modify-window", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			ModifyWindow: true,
		})
		assert.Contains(t, args, "--modify-window")
	})

	t.Run("--temp-dir", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			TempDir: "test",
		})
		assert.Contains(t, args, "--temp-dir", "test")
	})

	t.Run("--fuzzy", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Fuzzy: true,
		})
		assert.Contains(t, args, "--fuzzy")
	})

	t.Run("--compare-dest", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			CompareDest: "test",
		})
		assert.Contains(t, args, "--compare-dest", "test")
	})

	t.Run("--copy-dest=", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			CopyDest: "test",
		})
		assert.Contains(t, args, "--copy-dest", "test")
	})

	t.Run("--link-dest", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			LinkDest: "test",
		})
		assert.Contains(t, args, "--link-dest", "test")
	})

	t.Run("--compress", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Compress: true,
		})
		assert.Contains(t, args, "--compress")
	})

	t.Run("--compress-level", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			CompressLevel: 3,
		})
		assert.ElementsMatch(t, args, []string{"--compress-level", "3"})
	})

	t.Run("--skip-compress=", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			SkipCompress: []string{"1", "2"},
		})
		assert.ElementsMatch(t, args, []string{"--skip-compress", "1,2"})
	})

	t.Run("--cvs-exclude", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			CVSExclude: true,
		})
		assert.Contains(t, args, "--cvs-exclude")
	})

	t.Run("--stats", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Stats: true,
		})
		assert.Contains(t, args, "--stats")
	})

	t.Run("--human-readable", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			HumanReadable: true,
		})
		assert.Contains(t, args, "--human-readable")
	})

	t.Run("--progress", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Progress: true,
		})
		assert.Contains(t, args, "--progress")
	})

	t.Run("--info", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Info: "progress2",
		})
		assert.Contains(t, args, "--info")
		assert.Contains(t, args, "progress2")
	})

	t.Run("--exclude", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Exclude: []string{"foo", "bar", "\"baz\""},
		})
		assert.Contains(t, args, "--exclude=foo")
		assert.Contains(t, args, "--exclude=bar")
		assert.Contains(t, args, "--exclude=\"baz\"")
	})

	t.Run("--include", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Include: []string{"foo", "bar", "\"baz\""},
		})
		assert.Contains(t, args, "--include=foo")
		assert.Contains(t, args, "--include=bar")
		assert.Contains(t, args, "--include=\"baz\"")
	})

	t.Run("--filter", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Filter: "merge_filter.txt",
		})
		assert.Contains(t, args, "--filter=merge_filter.txt")
	})

	t.Run("--chown", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			Chown: "nobody:nobody",
		})
		assert.Contains(t, args, "--chown=nobody:nobody")
	})

	t.Run("--ipv4", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			IPv4: true,
		})
		assert.Contains(t, args, "--ipv4")
	})

	t.Run("--ipv6", func(t *testing.T) {
		args := GetArguments(RsyncOptions{
			IPv6: true,
		})
		assert.Contains(t, args, "--ipv6")
	})
}
