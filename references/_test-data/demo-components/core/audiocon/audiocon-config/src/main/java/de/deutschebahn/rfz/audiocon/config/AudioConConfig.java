package de.deutschebahn.rfz.audiocon.config;

/**
 * Demo configuration class for AudioCon.
 * This is a minimal example for build testing.
 */
public class AudioConConfig {

    private String environment;
    private int port;

    public AudioConConfig() {
        this.environment = "dev";
        this.port = 8080;
    }

    public String getEnvironment() {
        return environment;
    }

    public void setEnvironment(String environment) {
        this.environment = environment;
    }

    public int getPort() {
        return port;
    }

    public void setPort(int port) {
        this.port = port;
    }
}
