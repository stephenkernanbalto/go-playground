module go_playground/hello

go 1.17

replace go_playground/greetings => ../greetings

replace go_playground/greetings_with_error => ../greetings-with-error

replace go_playground/greetings_with_slices => ../greetings_with_slices

replace go_playground/greetings_multiple => ../greetings_multiple

require go_playground/greetings_multiple v0.0.0-00010101000000-000000000000
