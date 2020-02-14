-- This migration allows a faux devlocal certicate to have access to the Orders API locally.
-- Air Force devlocal certificate config/tls/devlocal-faux-air-force-orders.cer
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
	'618b8015-984e-419d-949f-c4957ec982e2',
	'e200f56764d66dc9e5e83919aec8efec8147fd6031e3d4f98a6199161a8d9918',
	'CN=localhost,OU=Not Air Force Orders,O=Not Air Force,L=Washington,ST=DC,C=US',
	true,
	true,
	true,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	now(),
	now());