import { useFetchJSON } from "../utils/hooks";
import { FileInfo } from "../datatypes";
import FilePreview from "../components/FilePreview";
import FolderView from "../components/FolderView";

interface FilesPageProps {
    params: { filepath: string };
}

export default function FilesPage({ params }: FilesPageProps) {
    const { filepath } = params;
    const fileInfo = useFetchJSON<FileInfo>(`/api/file/${filepath}`);

    if (fileInfo) {
        if (fileInfo.type === "dir" || fileInfo.type === "sharespace") {
            return <FolderView filePath={filepath} />
        } else {
            return <FilePreview filePath={filepath} fileInfo={fileInfo} />
        }
    } else {
        return null;
    }
}


