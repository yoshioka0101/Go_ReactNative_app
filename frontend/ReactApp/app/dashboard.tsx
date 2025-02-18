import { useEffect, useState } from 'react';
import { View, Text, ActivityIndicator, Button, StyleSheet } from 'react-native';
import { useRouter } from 'expo-router';

export default function Dashboard() {
  const router = useRouter();
  const [token, setToken] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchToken = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/auth/me", {
          method: "GET",
          credentials: "include", // Cookie を送信
        });
        const data = await response.json();
        if (data.token) {
          setToken(data.token);
        } else {
          router.replace("/");
        }
      } catch (error) {
        console.error("Error fetching token:", error);
        router.replace("/");
      } finally {
        setLoading(false);
      }
    };

    fetchToken();
  }, []);

  const handleLogout = async () => {
    await fetch("http://localhost:8080/api/auth/logout", { method: "POST", credentials: "include" });
    setToken(null);
    alert("ログアウトしました！！！！");
    router.replace("/");
  };

  if (loading) return <ActivityIndicator size="large" />;
  if (!token) return <Text>認証エラー: ログインしてください</Text>;

  return (
    <View>
      <Text>ログイン成功！</Text>
      <View style={styles.buttonContainer}>
        <Button title="ログアウト" onPress={handleLogout} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  buttonContainer: {
    marginTop: 16,
    alignItems: 'center',
  },
});
