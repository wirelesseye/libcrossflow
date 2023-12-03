import { FileList, FileListItem } from "../components/FileList";
import { useFetchJSON } from "../utils/hooks";
import { BoxIcon } from "lucide-react";

export default function SharespacesPage() {
    const sharespaces = useFetchJSON<string[]>("/api/file/sharespaces");

    return (
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
    );
}
