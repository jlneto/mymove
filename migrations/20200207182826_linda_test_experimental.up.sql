INSERT INTO "public"."users" ("id", "login_gov_uuid", "login_gov_email", "created_at", "updated_at", "deactivated", "active") VALUES ('0de0ee30-baa6-4a63-825c-8c59035300ed', '306d83aa-dea2-4d78-9ef5-c2c385300308', 'first.last@login.gov.test', '2020-02-07 18:30:30.932807', '2020-02-07 18:30:30.93281', 'f', 't');

INSERT INTO "public"."customers" ("id", "created_at", "updated_at", "user_id", "dod_id", "first_name", "last_name", "agency", "email", "phone")
	VALUES ('6ac40a00-e762-4f5f-b08d-3ea72a8e4b63', '2020-02-07 18:30:30.943493+00', '2020-02-07 18:30:30.943494+00', '0de0ee30-baa6-4a63-825c-8c59035300ed', '9952033988', 'Bqwehh', 'Bqqwehhh', 'AIR_FORCE', 'BQWEHHH@mail.com', '212-123-456');

INSERT INTO "public"."entitlements" ("id", "dependents_authorized", "total_dependents", "non_temporary_storage", "privately_owned_vehicle", "storage_in_transit", "created_at", "updated_at", "authorized_weight") VALUES ('f7473ed8-1b12-4174-9d4e-a8a1a0f7d9cb', 't', '1', 't', 't', '2', '2020-02-07 19:11:59.421065+00', '2020-02-07 19:11:59.421068+00', NULL);

INSERT INTO "public"."move_orders" ("id", "customer_id", "origin_duty_station_id", "destination_duty_station_id", "entitlement_id", "created_at", "updated_at", "confirmation_number", "order_number", "grade", "order_type", "order_type_detail", "date_issued", "report_by_date")
	VALUES ('6fca843a-a87e-4752-b454-0fac67aa4988', '6ac40a00-e762-4f5f-b08d-3ea72a8e4b63', '2d5ada83-e09a-47f8-8de6-83ec51694a86', 'd13070ed-3c29-4e13-a075-dc13606f7f69', 'f7473ed8-1b12-4174-9d4e-a8a1a0f7d9cb', '2020-02-07 18:30:31.005605+00', '2020-02-07 18:30:31.005607+00', '4CPY8P', NULL, 'E_7', 'GHC', 'TBD', '2020-01-15', '2020-02-15');

INSERT INTO "public"."move_task_orders" ("id", "move_order_id", "reference_id", "is_available_to_prime", "is_canceled", "created_at", "updated_at")
    VALUES ('5d4b25bb-eb04-4c03-9a81-ee0398cb779e', '6fca843a-a87e-4752-b454-0fac67aa4988', NULL, 'f', 'f', '2020-02-07 18:30:31.102009+00', '2020-02-07 18:30:31.10201+00');

INSERT INTO "public"."mto_shipments" ("id", "move_task_order_id", "scheduled_pickup_date", "requested_pickup_date", "customer_remarks", "pickup_address_id", "destination_address_id", "secondary_pickup_address_id", "secondary_delivery_address_id", "prime_estimated_weight", "prime_estimated_weight_recorded_date", "prime_actual_weight", "created_at", "updated_at", "shipment_type", "status")
    VALUES ('475579d5-aaa4-4755-8c43-c510381ff9b5', '5d4b25bb-eb04-4c03-9a81-ee0398cb779e', '2018-03-16', '2018-03-15', 'please treat gently', '71d3e3cb-856c-44e7-9cdf-516ee977ed1d', 'f2e0278a-2e52-4d72-941f-c80833a39060', NULL, NULL, '1000', '2018-03-20', '980', '2020-02-07 18:30:31.875872+00', '2020-02-07 18:30:31.875874+00', 'HHG', 'SUBMITTED');

INSERT INTO "public"."mto_service_items" ("id", "move_task_order_id", "mto_shipment_id", "re_service_id", "meta_id", "meta_type", "created_at", "updated_at")
    VALUES ('9db1bf43-0964-44ff-8384-3297951f6781', '5d4b25bb-eb04-4c03-9a81-ee0398cb779e', '475579d5-aaa4-4755-8c43-c510381ff9b5', '8d600f25-1def-422d-b159-617c7d59156e', '2b7c8614-fe42-4a92-9cf4-2bea2b5d2de9', 'unknown', '2020-02-07 18:30:31.883167+00', '2020-02-07 18:30:31.883169+00');

INSERT INTO "public"."mto_service_items" ("id", "move_task_order_id", "mto_shipment_id", "re_service_id", "meta_id", "meta_type", "created_at", "updated_at")
    VALUES ('d886431c-c357-46b7-a084-a0c85dd496d3', '5d4b25bb-eb04-4c03-9a81-ee0398cb779e', '475579d5-aaa4-4755-8c43-c510381ff9b5', '2bc3e5cb-adef-46b1-bde9-55570bfdd43e', 'b99a4740-3672-46c6-9064-f0cb73438bf2', 'unknown', '2020-02-07 18:30:31.890139+00', '2020-02-07 18:30:31.890141+00');