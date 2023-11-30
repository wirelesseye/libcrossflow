import { css } from "@emotion/css";
import { FileIcon } from "lucide-react";
import { FileInfo } from "../datatypes";
import Button from "./Button";

interface FilePreviewProps {
    filePath: string;
    fileInfo: FileInfo;
}

export default function FilePreview({ filePath: filepath, fileInfo: fileinfo }: FilePreviewProps) {
    return (
        <div className={styles.root}>
            {fileinfo ? (
                <div className={styles.inner}>
                    <div className={styles.iconContainer}>
                        <FileIcon
                            size={128}
                            absoluteStrokeWidth
                            strokeWidth={3}
                        />
                    </div>
                    <div className={styles.info}>
                        <div className={styles.filename}>{fileinfo.name}</div>
                        <a className={styles.download} href={`/api/download/${filepath}`} download="">
                            <Button>Download</Button>
                        </a>
                    </div>
                </div>
            ) : null}
        </div>
    );
}

const styles = {
    root: css`
        display: flex;
        border: 1px solid rgba(0, 0, 0, 0.1);
        border-radius: 15px;
        padding: 15px;
    `,
    inner: css`
        display: flex;
        margin: 0 auto;
        gap: 30px;
        flex-grow: 1;
        max-width: 800px;
        align-items: center;

        @media (max-width: 600px) {
            flex-direction: column;
        }
    `,
    iconContainer: css`
        display: flex;
        align-items: center;
        justify-content: center;
        width: 256px;
        height: 256px;
        border: 1px solid rgba(0, 0, 0, 0.1);
        border-radius: 15px;
    `,
    info: css`
        display: flex;
        flex-direction: column;
        gap: 15px;
        @media (max-width: 600px) {
            align-items: center;
        }
    `,
    filename: css`
        font-size: 1.5em;
    `,
    download: css`
        margin-top: auto;
    `
};
