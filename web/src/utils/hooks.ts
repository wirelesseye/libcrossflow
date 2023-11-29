import { useEffect, useMemo, useState } from "react";

export function useFetchJSON<T>(url: string) {
    const realUrl = useMemo(() => {
        if (import.meta.env.DEV) {
            return new URL(url, "http://localhost:4331").href;
        } else {
            return url;
        }
    }, [url])

    const [data, setData] = useState<T | null>(null);

    useEffect(() => {
        setData(null);
        fetch(realUrl)
            .then((data) => data.json())
            .then((json) => setData(json));
    }, [realUrl]);

    return data;
}
