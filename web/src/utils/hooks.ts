import { useEffect, useState } from "react";

export function useFetchJSON<T>(url: string) {
    let realUrl = url;
    if (import.meta.env.DEV) {
        realUrl = new URL(url, "http://localhost:4331").href;
    }

    const [data, setData] = useState<T | null>(null);

    useEffect(() => {
        fetch(realUrl)
            .then((data) => data.json())
            .then((json) => setData(json));
    }, []);

    return data;
}
