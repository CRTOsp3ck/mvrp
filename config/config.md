# ACHTUNG: ALWAYS KEEP THE CONFIG UP-TO-DATE
Until i find a way to deduplicate the necessary datas, please keep the config up-to-date manually!

1 [DTO]
- DTOs should contain all the data necessary for the sequenced actions.
- Repo search function will not be generated if it has no corresponding SearchDTO.

2 [Sequences]
- These are functions that are called from the repo

3 [Processing]
- Ensure the referenced processing functions (independant & sequential) are created before codegen.
- Processing functions are to be created manually
- Processing functions params should match the DTO. NO! WAIT! It should take in the DTO!

# NOTE:
Perhaps i should make a seperate DTO for Get, Create, Update, Delete and Search?
Lol I did actually.