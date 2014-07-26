##Slices

Herein are general slice functions created for my own practical use. Many of these functions involve using maps (i.e. hashtables) to count occurrences and remove duplicates, which is faster than doing these things the non-hashtable way. In most cases maps are even faster than hardcoded select statements. Don't believe me? Test it.

**Slice2Map_uint**, **Slice2Map_string**

Converts a slice to a map for fast lookup. Instead of looping through the slice to check if a value exists, the map can be checked instead in the same way as *String2Map*.

**SliceUniqueRetain_uint**, **SliceUniqueRetain_string**

Very fast method for removing duplicates from a slice, using maps, preserves the order of the original slice. Returns a new slice from a new array with all duplicates removed.

**SliceUnique_uint**, **SliceUnique_string**

Even faster than the above but does not preserve the order.

**SliceIntersect_uint**, **SliceIntersect_string**

Returns a new slice that consists only of values that existed in all slices.

**SliceRemove_uint**, **SliceRemove_string**

Return a new slice that consists of the values in the first slice that don't exist in any subsequent slices, i.e. values that are unique to the first slice only.

**SliceSquish_uint**, **SliceSquish_string**

Returns a new slice having combined (squished) the values of all slices into one, with duplicates removed and given priority (moved higher up the list.) It's like ORing all the slices together.

