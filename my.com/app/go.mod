module my.com/app

go 1.14

require my.com/firstmod v0.0.0
require my.com/secondmod v0.0.0

replace my.com/firstmod v0.0.0 => ../firstmod/v2
replace my.com/secondmod v0.0.0 => ../secondmod
