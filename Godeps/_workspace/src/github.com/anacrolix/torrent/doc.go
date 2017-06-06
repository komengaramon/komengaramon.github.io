/*
Package torrent implements a torrent client. Goals include:
 * Configurable data storage, such as file, mmap, and piece-based.
 * Downloading on demand: torrent.Reader will request only the data required to
   satisfy Reads, which is ideal for streaming and torrentfs.

BitTorrent features implemented include:
 * Protocol obfuscation
 * DHT
 * uTP
 * PEX
 * Magnet
 * IP Blocklists
 * Some IPv6
 * UDP Trackers

ConfigDir

A Client has a configurable ConfigDir that defaults to $HOME/.config/torrent.
Torrent metainfo files are cached at $CONFIGDIR/torrents/$infohash.torrent.
Infohashes in $CONFIGDIR/banned_infohashes cannot be added to the Client. A
P2P Plaintext Format blocklist is loaded from a file at the location specified
by the environment variable TORRENT_BLOCKLIST_FILE if set. otherwise from
$CONFIGDIR/blocklist. If $CONFIGDIR/packed-blocklist exists, this is memory-
mapped as a packed IP blocklist, saving considerable memory.

*/
package torrent
