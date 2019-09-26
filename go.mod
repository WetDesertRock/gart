module github.com/wetdesertrock/gart

go 1.12

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/jaffee/commandeer v0.1.0
	github.com/spf13/pflag v1.0.3
	github.com/ungerik/go-cairo v0.0.0-20180910143756-ed3ace63553d
	github.com/wetdesertrock/flexiconfig v0.0.0-00010101000000-000000000000
	github.com/wetdesertrock/spatialhash v1.0.0
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036
	layeh.com/gopher-luar v1.0.7
)

replace github.com/wetdesertrock/flexiconfig => ../flexiconfig

replace github.com/wetdesertrock/spatialhash => ../spatialhash
