module.exports = [
    {
        username: "root",
        email: "root@iit.vn",
        firstName: "Quản",
        lastName: "Trị Viên",
        node_id: "http://kong:8001",
        admin: true,
        password: process.env.KONGA_ADMIN_PASSWORD,
    },
    {
        username: "monitor",
        email: "monitor@iit.vn",
        firstName: "Nguyễn",
        lastName: "Văn Guest",
        node_id: "http://kong:8001",
        admin: false,
        password: process.env.KONGA_MONITOR_PASSWORD,
    }
]