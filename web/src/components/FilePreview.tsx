import { css } from "@emotion/css";
import { FileIcon } from "lucide-react";
import { FileInfo } from "../datatypes";

interface FilePreviewProps {
    filePath: string;
    fileInfo: FileInfo;
}

export default function FilePreview({ filePath: filepath, fileInfo: fileinfo }: FilePreviewProps) {
    return (
        <div className={styles.root}>
            {fileinfo ? (
                <>
                    <div className={styles.iconContainer}>
                        <FileIcon
                            size={128}
                            absoluteStrokeWidth
                            strokeWidth={3}
                        />
                    </div>
                    <div>
                        <div>{fileinfo.name}</div>
                        <a href={`/api/download/${filepath}`} download="">
                            Download
                        </a>
                    </div>
                </>
            ) : null}
        </div>
    );
}

const styles = {
    root: css`
        display: flex;
    `,
    iconContainer: css`
        display: flex;
        align-items: center;
        justify-content: center;
        width: 256px;
        height: 256px;
        border: 1px solid rgba(0, 0, 0, 0.2);
    `,
};
