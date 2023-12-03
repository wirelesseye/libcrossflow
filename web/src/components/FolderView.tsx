import { FileIcon, FolderIcon } from "lucide-react";
import { FileStat } from "../datatypes";
import { useFetchJSON } from "../utils/hooks";
import { FileList, FileListItem } from "./FileList";
import path from "../utils/path";

export interface FolderViewProps {
    filePath: string;
}

export default function FolderView({ filePath }: FolderViewProps) {
    const files = useFetchJSON<FileStat[]>(`/api/file/list/${filePath}`);

    return (
        <FileList>
            {files !== null
                ? files.map((file) => (
                      <FileListItem
                          key={file.name}
                          name={file.name}
                          href={path.join("/files/", filePath, file.name)}
                          icon={
                              file.type === "dir" ? (
                                  <FolderIcon size={20} />
                              ) : (
                                  <FileIcon size={20} />
                              )
                          }
                      />
                  ))
                : null}
        </FileList>
    );
}
