# Neko

Neko is a self-hosted drive alternative. The goal is for it to be able to store files on my server in the filesystem that I can access anywhere. 

What makes neko a finished and usable app for me:
- [ ] Multiple Users (need to be able to have other people store stuff on the server)
- [ ] Filesystem support (data is stored in a directory I specify via env)
- [ ] Clean UI (something that anyone can intuitively use)
- [ ] Low Memory Usage (no handling files in the server memory. data is written directly to disc and streamed from the disc. this needs to be lightweight)
- More stuff as and when I think of it

Future cool to have:
- S3 support so that i can do easier backups (thinking of rustfs or seaweedfs locally and cloudflare r2 as cloud backup)
- Mobile app
- File processing so that I can:
  - chat to find data with some llm
  - search for text in files from the UI
  - build a pipeline to have an llm describe images and videos (interval-based frame recognition) so that I can search for this similar to how i'd do for docs (this is gonna be fun to build)

Neko is essentially a ground-up re-write of a part of curiositi where I want to actually focus and build stuff properly instead of having it turn into another playground for me to build the new cool thing that i find interesting directly into this. I want to be structured about how I build here.
