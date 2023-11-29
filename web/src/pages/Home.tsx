import { css } from "@emotion/css";
import { FileList, FileListItem } from "../components/FileList";
import { useFetchJSON } from "../utils/hooks";
import { BoxIcon } from "lucide-react";

export default function HomePage() {
    const sharespaces = useFetchJSON<string[]>("/api/sharespaces");

    return (
        <div className={styles.root}>
            <FileList>
                {sharespaces
                    ? sharespaces.map((sharespace) => (
                          <FileListItem
                              key={sharespace}
                              icon={<BoxIcon size={20} />}
                              name={sharespace}
                              href={`/files/${sharespace}`}
                          />
                      ))
                    : null}
            </FileList>
        </div>
    );
}

const styles = {
    root: css`
        display: flex;
        flex-direction: column;
    `,
};
