module.exports = [
    {
        name: "Main gateway",
        type: "default",
        kong_admin_url: "http://kong:8001",
        health_checks: true,
        active: true,
    }
]