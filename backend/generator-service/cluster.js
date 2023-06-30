import Cluster from "puppeteer-cluster"

let cluster = null;

async function getCluster() {
    if (!cluster) {
        cluster = await Cluster.Cluster.launch({
            puppeteerOptions: { args: ["--no-sandbox","--font-render-hinting=none","--enable-font-antialiasing", "--disable-dev-shm-usage", "--disable-setuid-sandbox"] },
            concurrency: Cluster.Cluster.CONCURRENCY_CONTEXT,
            maxConcurrency: 2,
            headless: true,
        });
    }
    await cluster.task(async ({ page, data: dat }) => {
        await page.setContent(dat);
        await page.evaluateHandle('document.fonts.ready');
        const screen = await page.pdf({
            printBackground: true,
            format: "A4",
            margin: {
                top: "0",
                bottom: "0",
                left: "0",
                right: "0",
            },
        });
        return screen;
    });

    return cluster;
}


export { getCluster };