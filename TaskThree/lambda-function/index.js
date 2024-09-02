const { Client } = require('pg'); // PostgreSQL client library
const { DB_HOST, DB_NAME, DB_USER, DB_PASSWORD } = process.env;

exports.handler = async (event) => {
    const id = event.id;
    const ipAddress = event.ip_address;
    const isUpdate = event.is_update || false;

    if (!id || !ipAddress) {
        return {
            statusCode: 400,
            body: JSON.stringify({ success: false, message: 'ID and IP address are required' }),
        };
    }

    const client = new Client({
        host: DB_HOST,
        database: DB_NAME,
        user: DB_USER,
        password: DB_PASSWORD,
    });

    try {
        await client.connect();

        if (isUpdate) {
            // Update existing record
            const res = await client.query('UPDATE registrations SET ip_address = $1 WHERE id = $2', [ipAddress, id]);
            if (res.rowCount > 0) {
                return {
                    statusCode: 200,
                    body: JSON.stringify({ success: true }),
                };
            } else {
                return {
                    statusCode: 404,
                    body: JSON.stringify({ success: false, message: 'ID not found' }),
                };
            }
        } else {
            // Check if the record exists
            const res = await client.query('SELECT id FROM registrations WHERE id = $1', [id]);
            if (res.rows.length > 0) {
                return {
                    statusCode: 409,
                    body: JSON.stringify({ success: false, message: 'ID already exists' }),
                };
            }
            // Insert new record
            await client.query('INSERT INTO registrations (id, ip_address) VALUES ($1, $2)', [id, ipAddress]);
            return {
                statusCode: 200,
                body: JSON.stringify({ success: true }),
            };
        }
    } catch (error) {
        return {
            statusCode: 500,
            body: JSON.stringify({ success: false, message: error.message }),
        };
    } finally {
        await client.end();
    }
};
