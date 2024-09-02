const { Client } = require('pg'); // PostgreSQL client library
const axios = require('axios'); // HTTP client for making requests
const crypto = require('crypto'); // For generating new secret keys
const fs = require('fs'); // For file operations
const { DB_HOST, DB_NAME, DB_USER, DB_PASSWORD } = process.env;

exports.handler = async () => {
    const client = new Client({
        host: DB_HOST,
        database: DB_NAME,
        user: DB_USER,
        password: DB_PASSWORD,
    });

    try {
        await client.connect();

        // Get all registered EC2 instances
        const res = await client.query('SELECT id, ip_address, secret_key FROM registrations');
        const registrations = res.rows;

        for (let reg of registrations) {
            const { id, ip_address: ipAddress, secret_key: currentSecretKey } = reg;

            // Generate a new secret key
            const newSecretKey = crypto.randomBytes(16).toString('hex');

            // Update the secret key in the RDS database
            await client.query('UPDATE registrations SET secret_key = $1 WHERE id = $2', [newSecretKey, id]);

            // Update the secret key on the EC2 instance
            await axios.post(`http://${ipAddress}/update-secret-key`, {
                old_secret_key: currentSecretKey,
                new_secret_key: newSecretKey
            });

            // Optionally update the secret key file on the EC2 instance
            fs.writeFileSync('/path/to/secret_key_file', newSecretKey, 'utf8');
        }

        return {
            statusCode: 200,
            body: JSON.stringify({ success: true, message: 'Keys updated successfully' }),
        };
    } catch (error) {
        return {
            statusCode: 500,
            body: JSON.stringify({ success: false, message: error.message }),
        };
    } finally {
        await client.end();
    }
};
