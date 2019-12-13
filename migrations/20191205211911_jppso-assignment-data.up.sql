INSERT INTO jppso_regions VALUES ('df797b17-68c9-49d7-bc4d-e1c5e717826f', 'AGFM', 'JPPSO-Northeast', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('61908579-92b9-45bc-8651-ad48da299315', 'KKFA', 'JPPSO-North Central', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('2eabdded-be9e-4102-b656-5f5c1c8766c4', 'BGAC', 'JPPSO-Mid-Atlantic', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('ef49930a-81b7-4cc7-8afb-0cb0669b6ecc', 'CNNQ', 'JPPSO-Southeast', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('43d920b1-4c1b-4df3-862e-62725427eff8', 'HAFC', 'JPPSO-South Central', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('615673a2-d393-4d46-95f9-705fbeb6bc79', 'JEAT', 'JPPSO-Northwest', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('6448141f-7140-49f7-bbd8-c062f68435a8', 'LKNQ', 'JPPSO-Southwest', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('9138d6bc-3180-4163-a22a-69c366580024', 'LHNQ', 'USCG Base Alameda' , NOW(), NOW());
INSERT INTO jppso_regions VALUES ('81f99469-9604-439d-9978-b69bed5b178d', 'CLPK', 'USCG Base Miami Beach', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('560eb530-39d7-4363-8436-5ffb4bbe8e12', 'MAPS', 'USCG Base Kodiak', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('58510246-905b-4989-b16e-04806904134b', 'MAPK', 'USCG Base Ketchikan', NOW(), NOW());
INSERT INTO jppso_regions VALUES ('97611c2a-3ef6-4883-bf9a-3a7709f2acdd', 'MLNQ', 'JPPSO-Hawaii', NOW(), NOW());

INSERT INTO jppso_region_state_assignments VALUES ('1d3b0f1f-b6b1-4bb1-96d8-dbc72370c50b', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'HAFC'), 'Alabama', 'AL', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('3770b5af-6f0a-4ec5-a56d-713819b7055b', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'JEAT'), 'Alaska', 'AK', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('d91f6de0-73d5-4cba-a383-5a30f1f49647', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Arizona', 'AZ', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('c48a0707-e717-433c-9b68-24bed086aa2e', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'HAFC'), 'Arkansas', 'AR', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('bc26fd98-0fdb-4649-8ff1-71ccd5eefbf5', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'LKNQ'), 'California', 'CA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('76ee9db9-bee1-4417-beb5-720578c16c83', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Colorado', 'CO', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('59459f77-217f-413e-9dfc-b35f72df5834', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Connecticut', 'CT', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('03a18364-9ae1-47c3-9174-b97670facaae', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Delaware', 'DE', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('760cee55-e324-439a-bbbf-db18408324c1', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'BGAC'), 'District of Columbia', 'DC', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('2a4b3ac2-b9ed-4fc5-a5e6-dc77a02d5d6e', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'CNNQ'), 'Florida', 'FL', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('ad69da50-92d1-45a4-90a3-3caa027caa10', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'CNNQ'), 'Georgia', 'GA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('00a362f5-3aac-4b03-b2b0-09a48d72c056', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'MLNQ'), 'Hawaii', 'HI', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('a3d51575-f294-477c-ac31-6e9c9498522d', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'JEAT'), 'Idaho', 'ID', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('d3fc3b26-6e3d-4e82-b4e7-bb2a9601983a', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Illinois', 'IL', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('61a867e7-043c-45cc-95bd-1d7157c06bf6', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Indiana', 'IN', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('a164e0b2-5920-491f-9b32-8bd9c9e064fa', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Iowa', 'IA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('ff9a218c-c442-4e35-854c-92ecee652f82', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Kansas', 'KS', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('a8b9739c-c14f-4e7a-b319-a48d7aad46e8', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Kentucky', 'KY', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('06a9ffd3-fdf8-48fa-a245-e245301c8e14', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'HAFC'), 'Louisiana', 'LA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('8c16edad-6f5b-4710-befb-8eb5c5ded660', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Maine', 'ME', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('87228390-a460-4ef5-a270-ec931f4f087e', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'BGAC'), 'Maryland', 'MD', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('6b558cd9-c28c-4115-8e16-70d897f50d91', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Massachusetts', 'MA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('42798dde-4f8b-4645-abbf-ac589aee1d9c', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Michigan', 'MI', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('530b3026-be1f-4851-91d0-6eecdaf55589', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Minnesota', 'MN', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('c0a0f469-041b-442a-af29-d16cb167f807', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'HAFC'), 'Mississippi', 'MS', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('d7435c42-ae8b-4f63-9973-31179eb22dc5', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Missouri', 'MO', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('addd3d33-b413-41ed-bdad-e69de72bdab3', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Montana', 'MT', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('63a47702-18ee-4959-98e0-1e5aeb230e45', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Nebraska', 'NE', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('50778767-cd38-40ee-89e2-92a2624c2f57', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'LKNQ'), 'Nevada', 'NV', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('52f5e0d7-36b7-4d25-b638-b6bbf9d6f158', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'New Hampshire', 'NH', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('0d59d4b0-caef-4cc4-82f9-d60e94b5513f', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'New Jersey', 'NJ', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('1273070c-b08c-4078-b464-ce78a41fe83f', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'New Mexico', 'NM', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('cd4620c4-1c1b-4b0b-9923-ab794b51c4a7', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'New York', 'NY', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('132c4f30-baa5-41cd-8686-2e3adf8f26de', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'North Carolina', 'NC', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('5ac21af3-eab5-4de4-9bb1-c41f0b068bff', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'North Dakota', 'ND', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('3af061fe-5176-4251-aa95-1e2faed90c90', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Ohio', 'OH', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('5a14db89-7031-486f-9ab9-988df0a7a3fd', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Oklahoma', 'OK', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('81c8c3fe-1e30-4907-be63-5cb6efa22384', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'JEAT'), 'Oregon', 'OR', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('cf22a56e-7cd7-4652-bf80-9d7cbdbc2afb', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Pennsylvania', 'PA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('bbc1a150-9429-42af-91e0-1e50e0e9d07a', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Rhode Island', 'RI', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('04b1ff9d-ef88-4149-a586-9727bd8df1a7', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'CNNQ'), 'South Carolina', 'SC', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('362103fa-8520-499c-8c8e-06475a53063d', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'South Dakota', 'SD', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('99f8b6a8-385b-46b0-95f0-4aebf40a296a', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Tennessee', 'TN', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('50d66262-5c92-4cd3-9398-d978485cdf0b', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'HAFC'), 'Texas', 'TX', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('7c4d2d98-515a-4cfd-91a3-b72ed9711550', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Utah', 'UT', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('d74bead1-444c-45c0-bdc4-849d2e8da4b7', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'Vermont', 'VT', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('8896dbb0-31ae-4084-bde6-38e48532dab1', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'BGAC'), 'Virginia', 'VA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('47398d85-b931-4003-a89e-ae04e607b9ea', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'JEAT'), 'Washington', 'WA', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('9b37c23a-c59f-4d76-866e-834453b0050d', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'AGFM'), 'West Virginia', 'WV', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('d608be34-ff64-4d8a-b239-c915e78f60ee', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Wisconsin', 'WI', NOW(), NOW());
INSERT INTO jppso_region_state_assignments VALUES ('f136821b-8b34-46f7-8e3b-a6cce9608ba3', (SELECT id FROM jppso_regions WHERE jppso_regions.code = 'KKFA'), 'Wyoming', 'WY', NOW(), NOW());