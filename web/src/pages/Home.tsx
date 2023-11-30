import { useRedirect } from "../utils/router";

export default function HomePage() {
    useRedirect("/files");

    return null;
}
