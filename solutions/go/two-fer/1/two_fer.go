
// Package twofer has one function
// ShareWith(name string) 
// when called it will return 
//     - "One for name, one for me" if name is known
//     - "One for you, one for me" if name is empty
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	// if name is not empty
    if name != "" {
        return "One for " + name + ", one for me."
    }
	return "One for you, one for me."
}
