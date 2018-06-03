# AquaIllumination Profile manager

AI produce aquarium lighting equipment, which use a web interface to manage the lighting profile/schedule.

They allow you to upload/download profiles, or manually change them.

However if you've created a profile with many points, or acquired one from another reefer, arbitrary changes are time consuming to make.

This program allows you to affect the profile in various ways, saving time.

# Commands

## shift

Shift provides a mechanism to move an entire profile *left* or *right*, to suit your own time keeping.

  aip-manager -f ~/current.aip shift 0200

Shift works by taking the first time point in **any** of the spectrums, and moving it to the time specified.
It then takes that same difference and applies it to every other point in every spectrum.

It will error of that would move a non-zero point off the day in either direction.

In that case, move your end point first, and then shift the entire spectrum
