#include <string.h>

#include <gate.h>
#include <gate/service-inline.h>
#include <gate/services/clone.h>

static void clone_packet_received(struct gate_service *service, void *data, size_t size)
{
	if (size > sizeof (struct gate_service)) {
		size_t idlen = size - sizeof (struct gate_service);
		char id[idlen + 1];
		memcpy(id, data + sizeof (struct gate_service), idlen);
		id[idlen] = '\0';

		gate_debug("clone instance id: ");
		gate_debug(id);
		gate_debug("\n");
	} else {
		gate_debug("clone failed\n");
	}
}

static struct gate_service clone_service = {
	.name = CLONE_SERVICE_NAME,
	.received = clone_packet_received,
};

void main()
{
	gate_debug("clone program started\n");

	struct gate_service_registry *r = gate_service_registry_create();
	if (r == NULL)
		gate_exit(1);

	if (!gate_register_service(r, &clone_service))
		gate_exit(1);

	if (!gate_discover_services(r))
		gate_exit(1);

	if (clone_service.code == 0) {
		gate_debug("clone service not found\n");
		gate_exit(1);
	}

	clone_send(clone_service.code);

	while (1)
		gate_recv_for_services(r, 0);

	gate_debug("clone program exiting\n");
}
