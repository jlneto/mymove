-- This migration allows a faux devlocal certicate to have access to the Orders API locally.
-- Marine Corps devlocal certificate config/tls/devlocal-faux-marine-corps-orders.cer
INSERT INTO public.client_certs (
	id,
	sha256_digest,
	subject,
	allow_orders_api,
	allow_air_force_orders_read,
	allow_air_force_orders_write,
	allow_army_orders_read,
	allow_army_orders_write,
	allow_coast_guard_orders_read,
	allow_coast_guard_orders_write,
	allow_marine_corps_orders_read,
	allow_marine_corps_orders_write,
	allow_navy_orders_read,
	allow_navy_orders_write,
	created_at,
	updated_at)
VALUES (
	'1e32a17e-e23b-4cb7-9d81-9efa028b93cd',
	'9d89b450e72c9244d8a7fc748b835228d2df96500cc1f83e960caa48a18b219f',
	'CN=localhost,OU=Not Marine Corps Orders,O=Not Marine Corps,L=Washington,ST=DC,C=US',
	true,
	false,
	false,
	false,
	false,
	false,
	false,
	true,
	true,
	false,
	false,
	now(),
	now());