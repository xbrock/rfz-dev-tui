package de.deutschebahn.rfz.audiocon;

import de.deutschebahn.rfz.audiocon.config.AudioConConfig;

/**
 * Demo main class for AudioCon.
 */
public class AudioConApplication {

    public static void main(String[] args) {
        System.out.println("AudioCon starting...");
        AudioConConfig config = new AudioConConfig();
        System.out.println("Environment: " + config.getEnvironment());
        System.out.println("Port: " + config.getPort());
    }
}
