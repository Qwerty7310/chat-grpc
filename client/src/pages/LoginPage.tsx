import {useState} from "react"

export const LoginPage = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const onSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        console.log("LOGIN ATTEMPT:", {username, password});

        // TODO: вызов gRPC AuthService.Login()
    }

    return (
        <div style={styles.container}>
            <form style={styles.form} onSubmit={onSubmit}>
                <h2>Login</h2>

                <input
                    style={styles.input}
                    placeholder="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />

                <input
                    style={styles.input}
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />

                <button style={styles.button} type="submit">
                    Login
                </button>
            </form>
        </div>
    );
};

const styles: {[key: string]: React.CSSProperties} = {
    container: {
        width: "100%",
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        background: "#f5f5f5",
    },
    form: {
        display: "flex",
        flexDirection: "column",
        gap: "12px",
        padding: "20px",
        borderRadius: "8px",
        background: "#fff",
        width: "300px",
        boxShadow: "0 0 10px rgba(0,0,0,0.1)"
    },
    input: {
        padding: "10px",
        fontSize: "16px",
    },
    button: {
        padding: "10px",
        fontSize: "16px",
        cursor: "pointer",
    }
}