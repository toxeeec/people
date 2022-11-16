import { ActionIcon } from "@mantine/core";
import { IconArrowLeft } from "@tabler/icons";
import { useLocation, useNavigate } from "react-router-dom";

export default function LayoutGoBack() {
	const navigate = useNavigate();
	const { key } = useLocation();
	const handleClick = () => {
		key === "default" ? navigate("home") : navigate(-1);
	};
	return (
		<ActionIcon onClick={handleClick} radius="xl">
			<IconArrowLeft size={20} />
		</ActionIcon>
	);
}
