# grep
Implement a somewhat Perl-style grep in Go

    import "github.com/CharlesGriswold/grep"

    func Grep(ss []string, pat string) (ret []string)
    func Grepnot(ss []string, pat string) (ret []string)

`Grep` functions similarly to `grep /pat/, @string;` in Perl, and Grepnot is similar to `grep !/pat/, @string;`.
`Grep` returns the strings that match the pattern and `Grepnot` returns the strings that do not match the pattern.

An empty string does not match the pattern. The logic on this has to be explicit since Go's regular expressions don't
like empty strings.
