# oracle_trc

**WORK IN PROGRESS**

I wrote this program to extract sql queries from trc files generated on client side.

## trc_dump
Dump packets from trc files.

See cmd/trc_dump/readme.md for details

## queries
Dump SQL queries found in trc file

## Enabling trace files
Add following lines to SQLNET.ORA file

```
adr_base=off
TRACE_LEVEL_CLIENT = 16
TRACE_FILE_CLIENT = CLIENT
TRACE_DIRECTORY_CLIENT = d:\logs\oracle
```

## Do do

- [X] Display queries after a given date
- [X] Display executable associated with pid
- [X] Write tests independent from trace files (confidentiality) 
- [X] Sort outputs from several trace file in time order
- [ ] Understand binary format of packets (help wanted)
- [ ] Determine bind parameters value (help wanted)
- [ ] Decode responses (help wanted)

