#ifndef CN_SLOW_HASH_GO_H
#define CN_SLOW_HASH_GO_H

#include <unistd.h>
/*
 * @variant: https://github.com/monero-project/monero/commit/608fd6f14a6b9c0eeba2843fb14cbb235be0034f :with the intent to improve
 * Monero's resistance to ASICs and encourage mining decentralization.
 * @prehashed: do memcpy if hash already in memory
 */
void cn_slow_hash(const void * pptr, size_t dlen, char * h, int variant, int prehashed);

#endif
