SELECT
	incidents.id AS id,
	incidents.name as name,
	incidents.severity as severity,
	incidents.status as status,
	incident_types.id AS type_id,
	incident_types.name AS type_name
FROM
	incidents
LEFT JOIN
	incident_types ON incidents.type = incident_types.id