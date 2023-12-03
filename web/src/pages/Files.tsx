import { api, useFetchJSON } from "../utils/hooks";
import { FileStat } from "../datatypes";
import FilePreview from "../components/FilePreview";
import FolderView from "../components/FolderView";
import { ArrowUpFromLineIcon } from "lucide-react";
import { css } from "@emotion/css";
import { Button } from "../components/Button";
import path from "../utils/path";
import { Link } from "../components/Link";
import { Fragment, useCallback, useMemo } from "react";

interface FilesPageProps {
    params: { filepath: string };
}

export default function FilesPage({ params }: FilesPageProps) {
    const { filepath } = params;
    const fileStat = useFetchJSON<FileStat>(`/api/file/stat/${filepath}`);

    const onFileSelected = useCallback((e: Event) => {
        const fileInput = e.target as HTMLInputElement;

        if (fileInput.files) {
            const file = fileInput.files[0];
            const formData = new FormData();
            formData.append("file", file);
            formData.append("dirpath", filepath);
            fetch(api("/api/file/upload"), { method: "POST", body: formData });
        }
    }, []);

    const fileInput = useMemo(() => {
        const input = document.createElement("input");
        input.type = "file";
        input.addEventListener("change", onFileSelected);
        return input;
    }, []);

    const handleUpload = useCallback(() => {
        fileInput.click();
    }, []);

    return fileStat ? (
        <div>
            <div className={styles.header}>
                <div className={styles.path}>
                    {filepath.split("/").map((name, i) => (
                        <Fragment key={i}>
                            {i ? <div className={styles.sep}>/</div> : null}
                            <Link
                                tabIndex={-1}
                                href={path.join(
                                    "/files",
                                    ...filepath.split("/").slice(0, i + 1),
                                )}
                            >
                                <Button
                                    className={
                                        i === 0 ? styles.sharespace : undefined
                                    }
                                    variant="ghost"
                                >
                                    {name}
                                </Button>
                            </Link>
                        </Fragment>
                    ))}
                </div>
                {fileStat.type === "dir" ? (
                    <Button onClick={handleUpload}>
                        <ArrowUpFromLineIcon size={18} />
                    </Button>
                ) : null}
            </div>
            {fileStat.type === "dir" || fileStat.type === "sharespace" ? (
                <FolderView filePath={filepath} />
            ) : (
                <FilePreview filePath={filepath} fileInfo={fileStat} />
            )}
        </div>
    ) : null;
}

const styles = {
    header: css`
        display: flex;
        padding: 10px;
        overflow: auto;
        align-items: center;
    `,
    sep: css`
        display: flex;
        align-items: center;
        color: rgba(0, 0, 0, 0.2);
        user-select: none;
    `,
    path: css`
        display: flex;
        flex-grow: 1;
        white-space: nowrap;
        align-items: center;
    `,
    sharespace: css`
        font-weight: bold;
    `,
};
