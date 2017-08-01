#ifndef _GATE_SERVICES_CLONE_H
#define _GATE_SERVICES_CLONE_H

#include <stdint.h>

#include "../../gate.h"

#define CLONE_SERVICE_NAME "clone"

static inline void clone_send(uint16_t code)
{
	const struct gate_packet packet = {
		.size = sizeof (packet),
		.code = code,
	};

	gate_send_packet(&packet);
}

#endif
