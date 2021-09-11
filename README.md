# go-fathom

A package to leverage go-consequences with fathom data

The fathom data is formmatted as csv. Each state has its own csv file. Each row in a state is organized by the Fid of a structure in the NSI. For each structure we store a series of frequency based depths (in meters) for fluvial and coastal/pluvial for current condition and a future condition. The logic to read the fathom data and convert it into feet is in this library in the hazard_provider package.

The project relies on go-consequences for much of the computation of consequences. The default occupancy types (and therefore damage functions) are overriden in this library to represent the damage functions from oliver's 2020 report from nature. The default foundation heights are replaced with survey data from the NSI survey tool in 2020. 
