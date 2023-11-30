import { useRedirect } from "../Router";

export default function HomePage() {
    useRedirect("/files");

    return null;
}
