# Memory

In the previous part of this series, the computer was introduced as a processing unit, which fetches its instructions from memory. In almost every case, a computer's memory is not a simple contiguous region; the Gameboy is no exception in this regard. Since the GameBoxy CPU can access 65,536 individual locations on its address bus, a "memory map" can be drawn of all the regions where the CPU has access.

前回のパートでは、コンピュータがプロセッシングユニットとして紹介された。それはメモリから命令をフェッチする。ほとんどのケースにおいて、コンピュータのメモリは単純な順序区域ではない。ゲームボーイもその例外ではない。ゲームボーイのCPUは、アドレスバスを通じて 65,536 のロケーションにアクセスする。”メモリマップ”はCPUがアクセスする全ての区域を示す。

<img src="https://github.com/a2-ito/gist_images/blob/master/02_go-gameboy/01_the_fetch-decode-execute_loop.png" width="320">

A more detailed look at the memory regions is as follows:
- [0000-3FFF] Cartridge ROM, bank 0: The first 16,384 bytes of the cartridge program are always available at this point in the memory map. Special circumstances apply:
- - [0000-00FF] BIOS: When the CPU starts up, PC starts at ``000h``, which is the start of the 256-byte Gameboy BIOS code. Once the BIOS has run, it is removed from the memory map, and this area of the cartridge rom becomes addressable.

