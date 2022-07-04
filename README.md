# jewel-search go
[our love will be timeless](https://www.youtube.com/watch?v=HxXsnFTVR6M)

 Path of Exile Timeless Jewel seed searching tool, only works on command line but exports results to seeds.txt.

Enter the notable size you want to search, enter the notables, enter the mod you want to search (full list in decoders), but without spaces and %'s like write "base_stun_threshold_reduction_+%" like this "basestunthresholdreduction" program will calculate and give you a seed in seeds.txt.


Runtime on 8 notable search: [509.766322ms](https://pastebinp.com/Hrmby2O4VxxgsKdtdsov7A#eSGAyoH3XdHobZ1S4FU211nTs3dK3izOItyX8uCRqy0=)
Initial startup takes a bit long about 4 seconds, as it is searching every file for the CSVs and decoder.

I almost forgot, special thanks to @xeske for compressing CSV files! 
https://github.com/xeske/timeless

todo:
-> keystones
-> added bonuses to passive skills in seeds radius, addition to main search
