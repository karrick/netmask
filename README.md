PACKAGE

package netmask
    import "github.com/karrick/netmask"

    Package netmask returns the shorthand netmask from an IPv4 netmask
    string.

FUNCTIONS

func ConvertNetmaskToCIDR(input string) (num int, err error)
    ConvertNetmaskToCIDR returns CIDR notation from an IPv4 netmask string.


